package example;

import java.io.ByteArrayInputStream;
import java.io.DataInputStream;
import java.io.IOException;

public class Main {

    public static long decryptWinPrice(String enc_price, String secretKey, String signKey) {
        byte[] bytes = Decrypter.decryptSafeWebStringToByte(enc_price, secretKey, signKey);
        DataInputStream dis = new DataInputStream(new ByteArrayInputStream(bytes));
        long value = 0;
        try {
            value = dis.readLong();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return value;
    }

    public static void main(String args[]) throws Exception {
        // 加密key
        String encKeyStr = "54MPDbYaVicfYPdjQOKsXfoq4mqVmKxS";
        // 签名key
        String signKeyStr = "czyr0wPXEEBT2ORprTjoNo7ZYqxkJiA4";
        // 加密后的价格为
        String price1 = "mrvD2lYBAABjMC4gJjNNVBVwlIbbjjpAgyIudg";

        long price2 = decryptWinPrice(price1,encKeyStr, signKeyStr);
        System.out.println("解密后的价格为：" + price2);
    }
}
