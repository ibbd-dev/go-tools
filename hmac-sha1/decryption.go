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
	"hash"
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

	// 加解密算法的new函数, 例如sha1.New
	algo_new func() hash.Hash
}

func (c *Crypto) SetNew(algo_new func() hash.Hash) {
	c.algo_new = algo_new
}

// {initialization_vector (16 bytes)}{ciphertext}{integrity (4 bytes)}
func (c *Crypto) Decrypt(ciphertext []byte) ([]byte, error) {
	if c.algo_new == nil {
		c.algo_new = sha1.New
	}

	// Step 1. find the length of initialization vector and clear text.
	ciphertext_len := len(ciphertext)
	ciphertext_end := ciphertext_len - SIGNATURE_SIZE
	plaintext_len := ciphertext_end - INITIALIZATION_VECTOR_SIZE
	if plaintext_len < 0 {
		return nil, errors.New("plaintext_len < 0")
	}

	iv := ciphertext[0:INITIALIZATION_VECTOR_SIZE]

	// Step 2. recover clear text
	var plaintext []byte = make([]byte, plaintext_len)
	add_iv_counter_byte := true

	var mac hash.Hash
	for ciphertext_begin, plaintext_begin := INITIALIZATION_VECTOR_SIZE, 0; ciphertext_begin < ciphertext_end; {
		mac = hmac.New(c.algo_new, c.EKey)
		mac.Write(iv)
		pad := mac.Sum(nil)

		for i := 0; i < BLOCK_SIZE && ciphertext_begin != ciphertext_end; i++ {
			plaintext[plaintext_begin] = byte(ciphertext[ciphertext_begin] ^ pad[i])
			plaintext_begin++
			ciphertext_begin++
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
	sign := ciphertext[ciphertext_end:ciphertext_len]
	computed_sign := c.calSign(plaintext, ciphertext)
	if hmac.Equal(computed_sign, sign) {
		return plaintext, nil
	}

	return nil, errors.New("computedSignature != signature")
}

// 计算签名
func (c *Crypto) calSign(plaintext, ciphertext []byte) []byte {
	mac := hmac.New(c.algo_new, c.IKey)
	mac.Write(plaintext)
	mac.Write(ciphertext[:INITIALIZATION_VECTOR_SIZE])
	mac_sign := mac.Sum(nil)

	return mac_sign[:SIGNATURE_SIZE]
}
