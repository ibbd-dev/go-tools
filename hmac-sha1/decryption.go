/**
 * Performs the decryption algorithm.
 *
 * This method decrypts the ciphertext using the encryption key and verifies
 * the integrity bits with the integrity key. The encrypted format is:
 * {initialization_vector (16 bytes)}{ciphertext}{integrity (4 bytes)}
 * https://developers.google.com/ad-exchange/rtb/response-guide/decrypt-
 * hyperlocal,
 * https://developers.google.com/ad-exchange/rtb/response-guide/decrypt
 * -price and https://support.google.com/adxbuyer/answer/3221407?hl=en have
 * more details about the encrypted format of hyperlocal, winning price,
 * IDFA, hashed IDFA and Android Advertiser ID.
 */
package hmacSHA1

import (
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	//"fmt"
)

const (
	/** The length of the initialization vector */
	INITIALIZATION_VECTOR_SIZE = 16

	/** The length of the signature */
	SIGNATURE_SIZE = 4

	/** The fixed block size for the block cipher */
	BLOCK_SIZE = 20
)

type Crypto struct {
	EKey []byte
	IKey []byte
}

// {initialization_vector (16 bytes)}{ciphertext}{integrity (4 bytes)}
func (c *Crypto) Decrypt(ciphertext []byte) ([]byte, error) {
	// Step 1. find the length of initialization vector and clear text.
	ciphertext_len := len(ciphertext)
	ciphertext_end := ciphertext_len - SIGNATURE_SIZE
	plaintext_len := ciphertext_end - INITIALIZATION_VECTOR_SIZE
	//fmt.Printf("len ciphertext_len = %d, plaintext_len = %d\n", ciphertext_len, plaintext_len)
	if plaintext_len < 0 {
		return nil, errors.New("plaintext_len < 0")
	}

	iv := ciphertext[0:INITIALIZATION_VECTOR_SIZE]

	// Step 2. recover clear text
	var plaintext []byte = make([]byte, plaintext_len)
	add_iv_counter_byte := true

	mac := hmac.New(sha1.New, c.EKey)
	for ciphertext_begin, plaintext_begin := INITIALIZATION_VECTOR_SIZE, 0; ciphertext_begin < ciphertext_end; {
		mac.Reset()
		pad := mac.Sum(iv)
		//fmt.Printf("%x\n", pad)

		for i := 0; i < BLOCK_SIZE && ciphertext_begin != ciphertext_end; i++ {
			plaintext[plaintext_begin] = byte(ciphertext[ciphertext_begin] ^ pad[i])
			plaintext_begin++
			ciphertext_begin++
			//fmt.Printf("%x\n", plaintext)
		}

		if !add_iv_counter_byte {
			index := len(iv) - 1
			iv[index]++
			add_iv_counter_byte = iv[index] == 0
		}

		if add_iv_counter_byte {
			add_iv_counter_byte = false
			iv = iv[0 : len(iv)+1]
		}
	}

	// Step 3. Compute integrity hash. The input to the HMAC is
	// clear_text
	// followed by initialization vector, which is stored in the 1st
	// section
	// or ciphertext.
	mac = hmac.New(sha1.New, c.IKey)
	mac.Write(plaintext)
	mac.Write(ciphertext[0:INITIALIZATION_VECTOR_SIZE])
	mac_sign := mac.Sum(nil)

	computedSignature := mac_sign[0:SIGNATURE_SIZE]
	signature := ciphertext[ciphertext_end:ciphertext_len]
	if !hmac.Equal(computedSignature, signature) {
		return plaintext, errors.New("computedSignature != signature")
	}

	return plaintext, nil
}

func (c *Crypto) Decrypt2(ciphertext []byte) ([]byte, error) {
	ciphertext_len := len(ciphertext)
	ciphertext_end := ciphertext_len - SIGNATURE_SIZE
	plaintext_len := ciphertext_end - INITIALIZATION_VECTOR_SIZE

	iv, ciphertext, sig := ciphertext[0:INITIALIZATION_VECTOR_SIZE], ciphertext[INITIALIZATION_VECTOR_SIZE:ciphertext_end], ciphertext[ciphertext_end:ciphertext_len]

	var plaintext []byte = make([]byte, plaintext_len)
	mac := hmac.New(sha1.New, c.EKey)
	pad := mac.Sum(iv)
	for i := 0; i < plaintext_len; i++ {
		plaintext[i] = byte(ciphertext[i] ^ pad[i])
	}

	mac = hmac.New(sha1.New, c.IKey)
	mac.Write(plaintext)
	conf_sig := mac.Sum(iv)

	if !hmac.Equal(sig, conf_sig) {
		return plaintext, errors.New("sig != conf_sig")
	}

	return plaintext, nil
}
