package contracts

// TD10ReportBody is an auto generated low-level Go binding around an user-defined struct.
type TD10ReportBody struct {
	TeeTcbSvn      [16]byte
	MrSeam         []byte
	MrsignerSeam   []byte
	SeamAttributes [8]byte
	TdAttributes   [8]byte
	XFAM           [8]byte
	MrTd           []byte
	MrConfigId     []byte
	MrOwner        []byte
	MrOwnerConfig  []byte
	RtMr0          []byte
	RtMr1          []byte
	RtMr2          []byte
	RtMr3          []byte
	ReportData     []byte
}