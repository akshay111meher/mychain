package crypto

import(
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/md5"
	"crypto/x509"
	"encoding/hex"
	"fmt"
    _"reflect"
	"hash"
	"math/big"
	"io"
)

func GenerateKeys() (privKeyBytes,pubKeyBytes[]byte){
	pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256

 	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

 	if err != nil {
 		fmt.Println(err)
 	}else{
		//  fmt.Println(reflect.TypeOf(privatekey))
		// fmt.Println(privatekey)
	}

	private_key_bytes, _ := x509.MarshalECPrivateKey(privatekey)
	public_key_bytes, _ := x509.MarshalPKIXPublicKey(&privatekey.PublicKey)
	// fmt.Println(*privatekey)
	//  fmt.Println(reflect.TypeOf(privatekey))

	//  var h hash.Hash
 	// h = md5.New()
 	// r := big.NewInt(0)
 	// s := big.NewInt(0)

 	// io.WriteString(h, "This is the message which has to be signed")
 	// signhash := h.Sum(nil)

 	// r, s, serr := ecdsa.Sign(rand.Reader, privatekey, signhash)
 	// if serr != nil {
 	// 	fmt.Println(err)
 	// }
	//  fmt.Printf("h : %x\n",h)
 	// signature := r.Bytes()
 	// signature = append(signature, s.Bytes()...)
	// fmt.Printf("Signhash : %x\n", signhash)
 	// fmt.Printf("Signature : %x\n", signature)

 	// // Verify
 	// verifystatus := ecdsa.Verify(&pubkey, signature, r, s)
 	// fmt.Println(verifystatus) // should be true
	return private_key_bytes,public_key_bytes
}

func GetSignature(data string, prk string) (string,string){
	privateKey,_ := hex.DecodeString(prk)
	privKey,_ := x509.ParseECPrivateKey(privateKey)
	
	var h hash.Hash
	h = md5.New()
	io.WriteString(h,data)
	signhash := h.Sum(nil)

	// fmt.Println(signhash)
	
	r := big.NewInt(0)
	s := big.NewInt(0) 
	r,s,err := ecdsa.Sign(rand.Reader, privKey, signhash)
	
	if err != nil {
		return hex.EncodeToString([]byte("")),hex.EncodeToString([]byte(""))
	}else{
		return hex.EncodeToString(r.Bytes()),hex.EncodeToString(s.Bytes())
	}
}

func Verify(data, puk,r,s string) bool{
	public_key_bytes,_ := hex.DecodeString(puk)
	public_key, err := x509.ParsePKIXPublicKey(public_key_bytes)
	
	if err != nil {
		return false
	}

	var h hash.Hash
	h = md5.New()
	io.WriteString(h,data)
	signhash := h.Sum(nil)
	// fmt.Println(signhash)
	
	rB,_ := hex.DecodeString(r)
	rS,_ := hex.DecodeString(s)
	r_big_int := big.NewInt(0)
	s_big_int := big.NewInt(0)	

	r_big_int = r_big_int.SetBytes(rB)
	s_big_int = s_big_int.SetBytes(rS)

	switch public_key := public_key.(type) {
		case *ecdsa.PublicKey:
			return ecdsa.Verify(public_key, signhash, r_big_int, s_big_int)
		default:
			return false
		}
}