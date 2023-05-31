package nfe

import (
	"encoding/xml"
	"time"
)

type NfeJsonList struct {
	Nfe []Nfe `json:"nfe"`
}

type Nfe struct {
	Data Data `json:"data"`
}

type Data struct {
	Xml string `json:"xml"`
}

type Xml struct {
	XMLName xml.Name `xml:"nfeProc"`
	NFe     NFe      `xml:"NFe"`
	ProtNFe ProtNFe  `xml:"protNFe"`
}

// type NfeProc struct {
// 	XMLName xml.Name `xml:"nfeProc"`
// 	NFe     NFe      `xml:"NFe"`
// }

type NFe struct {
	XMLName xml.Name `xml:"NFe"`
	InfNFe  InfNFe   `xml:"infNFe"`
}

type ProtNFe struct {
	XMLName xml.Name `xml:"protNFe"`
	InfProt InfProt  `xml:"infProt"`
}

type InfProt struct {
	XMLName  xml.Name `xml:"infProt"`
	ChaveNfe string   `xml:"chNFe"`
}

type InfNFe struct {
	Ide   Ide   `xml:"ide"`
	Det   Det   `xml:"det"`
	Total Total `xml:"total"`
}

type Ide struct {
	DataEmissao time.Time `xml:"dhEmi"`
}

type Det struct {
	Prod Prod `xml:"prod"`
}

type Prod struct {
	Cfop string `xml:"CFOP"`
}

type Total struct {
	ICMSTot ICMSTot `xml:"ICMSTot"`
}

type ICMSTot struct {
	ValorNfe string `xml:"vNF"`
}

// type Xml struct {
// 	ChaveNfe    string    `xml:"id"`
// 	ValorNfe    string    `xml:"vNF"`
// 	Cfop        string    `xml:"CFOP"`
// 	DataEmissao time.Time `xml:"dhEmi"`
// }
