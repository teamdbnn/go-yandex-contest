package models

type JSONNode struct {
	Array               bool   `json:"array,omitempty"`
	BigDecimal          bool   `json:"bigDecimal,omitempty"`
	BigInteger          bool   `json:"bigInteger,omitempty"`
	Binary              bool   `json:"binary,omitempty"`
	Boolean             bool   `json:"boolean,omitempty"`
	ContainerNode       bool   `json:"containerNode,omitempty"`
	Double              bool   `json:"double,omitempty"`
	Empty               bool   `json:"empty,omitempty"`
	Float               bool   `json:"float,omitempty"`
	FloatingPointNumber bool   `json:"floatingPointNumber,omitempty"`
	Int                 bool   `json:"int,omitempty"`
	IntegralNumber      bool   `json:"integralNumber,omitempty"`
	Long                bool   `json:"long,omitempty"`
	MissingNode         bool   `json:"missingNode,omitempty"`
	NodeType            string `json:"nodeType,omitempty"`
	Null                bool   `json:"null,omitempty"`
	Number              bool   `json:"number,omitempty"`
	Object              bool   `json:"object,omitempty"`
	Pojo                bool   `json:"pojo,omitempty"`
	Short               bool   `json:"short,omitempty"`
	Textual             bool   `json:"textual,omitempty"`
	ValueNode           bool   `json:"valueNode,omitempty"`
}
