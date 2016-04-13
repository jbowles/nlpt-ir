package nlptir

import (
	"bytes"
	"encoding/gob"
	//"fmt"
)

// TYPES //////////////////////////////////////////////////

/*
	var cache bytes.Buffer
	vstream := vf.EncodeVectorStream(cache)
	if vstream.EncodeError != nil {
		log.Fatal("encode:", vstream.EncodeError)
	}
	fmt.Printf("%+v\n", vstream)


	//   decoding way one: right off the VectorStream struct
	var vecField VecField
	res := vstream.GobDecoder.Decode(&vecField)
	fmt.Printf("%v, %v", vecField, res)

	//   decoding way two: pass the VectorStream struct as param
	resgob, reserr := DecodeVectorStream(vstream)
	fmt.Printf("%+v, %v", resgob, reserr)

	//   decoding way three: decode VectorStream.Bytes as []byte
	resgob, reserr := DecodeVectorStream(vstream)
	fmt.Printf("%+v, %v", resgob, reserr)
*/

type VectorStream struct {
	GobDecoder   *gob.Decoder
	EncodeError  error
	ByteEncoding []byte
}

func (vf VecField) EncodeVectorStream(byteCache *bytes.Buffer) *VectorStream {
	encErr := gob.NewEncoder(byteCache).Encode(vf)
	dec := gob.NewDecoder(byteCache)
	return &VectorStream{
		GobDecoder:   dec,
		EncodeError:  encErr,
		ByteEncoding: byteCache.Bytes(),
	}
}

func DecodeVectorStream(vs *VectorStream) (VecField, error) {
	var vecField VecField
	err := vs.GobDecoder.Decode(&vecField)
	return vecField, err
}

func DecodeVectorStreamBytes(vStreamBytes []byte) (VecField, error) {
	decBuf := bytes.NewBuffer(vStreamBytes)
	var vf VecField
	err := gob.NewDecoder(decBuf).Decode(&vf)
	return vf, err
}

/*

NOT WORKING YET WITH MAPS...
	vbyt, err := vf.MarshalVecField()
	res := vf.UnmarshalVecField(vbyt)
func (vf *VecField) MarshalVecField() ([]byte, error) {
	//encoding to text
	var b bytes.Buffer
	fmt.Fprintln(&b, vf.Space)
	return b.Bytes(), nil
}

func (vf *VecField) UnmarshalVecField(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &vf.Space)
	return err
}
*/
