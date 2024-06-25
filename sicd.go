package sicd

import (
	"encoding/xml"
	"time"
)

type SICD struct {
	XMLName         xml.Name        `xml:"SICD"`
	CollectionInfo  CollectionInfo  `xml:"CollectionInfo"`
	ImageCreation   *ImageCreation  `xml:"ImageCreation,omitempty"`
	ImageData       ImageData       `xml:"ImageData"`
	GeoData         GeoData         `xml:"GeoData"`
	Grid            Grid            `xml:"Grid"`
	Timeline        Timeline        `xml:"Timeline"`
	Position        Position        `xml:"Position"`
	RadarCollection RadarCollection `xml:"RadarCollection"`
	ImageFormation  ImageFormation  `xml:"ImageFormation"`
	SCPCOA          SCPCOA          `xml:"SCPCOA"`
	Radiometric     *Radiometric    `xml:"Radiometric,omitempty"`
	Antenna         *Antenna        `xml:"Antenna,omitempty"`
	ErrorStatistics *ErrorStatistics `xml:"ErrorStatistics,omitempty"`
	MatchInfo       *MatchInfo      `xml:"MatchInfo,omitempty"`
	RgAzComp        *RgAzComp       `xml:"RgAzComp,omitempty"`
	PFA             *PFA            `xml:"PFA,omitempty"`
	RMA             *RMA            `xml:"RMA,omitempty"`
}

type CollectionInfo struct {
	CollectorName string `xml:"CollectorName"`
	CoreName      string `xml:"CoreName"`
	CollectType   string `xml:"CollectType,omitempty"`
	RadarMode     struct {
		ModeType string `xml:"ModeType"`
		ModeID   string `xml:"ModeID,omitempty"`
	} `xml:"RadarMode"`
	Classification string `xml:"Classification"`
	CountryCode    []string `xml:"CountryCode,omitempty"`
}

type ImageCreation struct {
	Application string     `xml:"Application,omitempty"`
	DateTime    *time.Time `xml:"DateTime,omitempty"`
	Site        string     `xml:"Site,omitempty"`
	Profile     string     `xml:"Profile,omitempty"`
}

type ImageData struct {
	PixelType  string `xml:"PixelType"`
	NumRows    uint   `xml:"NumRows"`
	NumCols    uint   `xml:"NumCols"`
	FirstRow   uint   `xml:"FirstRow"`
	FirstCol   uint   `xml:"FirstCol"`
	FullImage  struct {
		NumRows uint `xml:"NumRows"`
		NumCols uint `xml:"NumCols"`
	} `xml:"FullImage"`
	SCPPixel RowCol `xml:"SCPPixel"`
}

type RowCol struct {
	Row int `xml:"Row"`
	Col int `xml:"Col"`
}

type GeoData struct {
	EarthModel string `xml:"EarthModel"`
	SCP        struct {
		ECF XYZType               `xml:"ECF"`
		LLH LatLonHeightRestrType `xml:"LLH"`
	} `xml:"SCP"`
	ImageCorners struct {
		ICP []LatLonCornerStringType `xml:"ICP"`
	} `xml:"ImageCorners"`
}

type XYZType struct {
	X float64 `xml:"X"`
	Y float64 `xml:"Y"`
	Z float64 `xml:"Z"`
}

type LatLonHeightRestrType struct {
	Lat  float64 `xml:"Lat"`
	Lon  float64 `xml:"Lon"`
	HAE  float64 `xml:"HAE"`
}

type LatLonCornerStringType struct {
	Lat   float64 `xml:"Lat"`
	Lon   float64 `xml:"Lon"`
	Index string  `xml:"index,attr"`
}

// Grid represents the Grid element
type Grid struct {
	ImagePlane string    `xml:"ImagePlane"`
	Type       string    `xml:"Type"`
	TimeCOAPoly Poly2DType `xml:"TimeCOAPoly"`
	Row        DirParamType `xml:"Row"`
	Col        DirParamType `xml:"Col"`
}

// DirParamType represents the Row and Col elements in Grid
type DirParamType struct {
	UVectECF   XYZType   `xml:"UVectECF"`
	SS         float64   `xml:"SS"`
	ImpRespWid float64   `xml:"ImpRespWid"`
	Sgn        int       `xml:"Sgn"`
	ImpRespBW  float64   `xml:"ImpRespBW"`
	KCtr       float64   `xml:"KCtr"`
	DeltaK1    float64   `xml:"DeltaK1"`
	DeltaK2    float64   `xml:"DeltaK2"`
	DeltaKCOAPoly *Poly2DType `xml:"DeltaKCOAPoly,omitempty"`
	WgtType    *struct {
		WindowName string      `xml:"WindowName"`
		Parameter  []Parameter `xml:"Parameter,omitempty"`
	} `xml:"WgtType,omitempty"`
	WgtFunct   *struct {
		Wgt  []struct {
			Value float64 `xml:",chardata"`
			Index int     `xml:"index,attr"`
		} `xml:"Wgt"`
		Size int `xml:"size,attr"`
	} `xml:"WgtFunct,omitempty"`
}

// Timeline represents the Timeline element
type Timeline struct {
	CollectStart    time.Time `xml:"CollectStart"`
	CollectDuration float64   `xml:"CollectDuration"`
	IPP             *struct {
		Set  []IPPSet `xml:"Set"`
		Size int      `xml:"size,attr"`
	} `xml:"IPP,omitempty"`
}

// IPPSet represents the Set element in IPP
type IPPSet struct {
	TStart   float64    `xml:"TStart"`
	TEnd     float64    `xml:"TEnd"`
	IPPStart int        `xml:"IPPStart"`
	IPPEnd   int        `xml:"IPPEnd"`
	IPPPoly  Poly1DType `xml:"IPPPoly"`
	Index    int        `xml:"index,attr"`
}

// Position represents the Position element
type Position struct {
	ARPPoly   XYZPolyType    `xml:"ARPPoly"`
	GRPPoly   *XYZPolyType   `xml:"GRPPoly,omitempty"`
	TxAPCPoly *XYZPolyType   `xml:"TxAPCPoly,omitempty"`
	RcvAPC    *RcvAPCType    `xml:"RcvAPC,omitempty"`
}

// RcvAPCType represents the RcvAPC element in Position
type RcvAPCType struct {
	RcvAPCPoly []XYZPolyAttributeType `xml:"RcvAPCPoly"`
	Size       int                    `xml:"size,attr"`
}

// RadarCollection represents the RadarCollection element
type RadarCollection struct {
	TxFrequency    TxFrequencyType     `xml:"TxFrequency"`
	RefFreqIndex   *int                `xml:"RefFreqIndex,omitempty"`
	Waveform       *WaveformType       `xml:"Waveform,omitempty"`
	TxPolarization Polarization1Type   `xml:"TxPolarization"`
	TxSequence     *TxSequenceType     `xml:"TxSequence,omitempty"`
	RcvChannels    RcvChannelsType     `xml:"RcvChannels"`
	Area           *AreaType           `xml:"Area,omitempty"`
	Parameter      []Parameter         `xml:"Parameter,omitempty"`
}

// TxFrequencyType represents the TxFrequency element in RadarCollection
type TxFrequencyType struct {
	Min float64 `xml:"Min"`
	Max float64 `xml:"Max"`
}

// WaveformType represents the Waveform element in RadarCollection
type WaveformType struct {
	WFParameters []WFParametersType `xml:"WFParameters"`
	Size         int                `xml:"size,attr"`
}

// WFParametersType represents the WFParameters element in Waveform
type WFParametersType struct {
	TxPulseLength   *float64 `xml:"TxPulseLength,omitempty"`
	TxRFBandwidth   *float64 `xml:"TxRFBandwidth,omitempty"`
	TxFreqStart     *float64 `xml:"TxFreqStart,omitempty"`
	TxFMRate        *float64 `xml:"TxFMRate,omitempty"`
	RcvDemodType    *string  `xml:"RcvDemodType,omitempty"`
	RcvWindowLength *float64 `xml:"RcvWindowLength,omitempty"`
	ADCSampleRate   *float64 `xml:"ADCSampleRate,omitempty"`
	RcvIFBandwidth  *float64 `xml:"RcvIFBandwidth,omitempty"`
	RcvFreqStart    *float64 `xml:"RcvFreqStart,omitempty"`
	RcvFMRate       *float64 `xml:"RcvFMRate,omitempty"`
	Index           int      `xml:"index,attr"`
}

// Continuing with more struct definitions...

type TxSequenceType struct {
	TxStep []TxStepType `xml:"TxStep"`
	Size   int          `xml:"size,attr"`
}

type TxStepType struct {
	WFIndex        *int             `xml:"WFIndex,omitempty"`
	TxPolarization *Polarization2Type `xml:"TxPolarization,omitempty"`
	Index          int              `xml:"index,attr"`
}

type RcvChannelsType struct {
	ChanParameters []ChanParametersType `xml:"ChanParameters"`
	Size           int                  `xml:"size,attr"`
}

type ChanParametersType struct {
	TxRcvPolarization DualPolarizationType `xml:"TxRcvPolarization"`
	RcvAPCIndex       *int                 `xml:"RcvAPCIndex,omitempty"`
	Index             int                  `xml:"index,attr"`
}

type AreaType struct {
	Corner CornerType `xml:"Corner"`
	Plane  *PlaneType `xml:"Plane,omitempty"`
}

type CornerType struct {
	ACP []LatLonHAECornerRestrType `xml:"ACP"`
}

type PlaneType struct {
	RefPt       RefPtType       `xml:"RefPt"`
	XDir        XYDirType       `xml:"XDir"`
	YDir        XYDirType       `xml:"YDir"`
	SegmentList *SegmentListType `xml:"SegmentList,omitempty"`
	Orientation *OrientationType `xml:"Orientation,omitempty"`
}

type RefPtType struct {
	ECF   XYZType `xml:"ECF"`
	Line  float64 `xml:"Line"`
	Sample float64 `xml:"Sample"`
	Name  string  `xml:"name,attr"`
}

type XYDirType struct {
	UVectECF      XYZType `xml:"UVectECF"`
	LineSpacing   float64 `xml:"LineSpacing"`
	NumLines      int     `xml:"NumLines"`
	FirstLine     int     `xml:"FirstLine"`
}

type SegmentListType struct {
	Segment []SegmentType `xml:"Segment"`
	Size    int           `xml:"size,attr"`
}

type SegmentType struct {
	StartLine   int    `xml:"StartLine"`
	StartSample int    `xml:"StartSample"`
	EndLine     int    `xml:"EndLine"`
	EndSample   int    `xml:"EndSample"`
	Identifier  string `xml:"Identifier"`
	Index       int    `xml:"index,attr"`
}

// Add more struct definitions for other complex types...

type Poly1DType struct {
	Coef  []PolyCoef1DType `xml:"Coef"`
	Order1 int              `xml:"order1,attr"`
}

type Poly2DType struct {
	Coef   []PolyCoef2DType `xml:"Coef"`
	Order1 int              `xml:"order1,attr"`
	Order2 int              `xml:"order2,attr"`
}

type PolyCoef1DType struct {
	Coef     float64 `xml:",chardata"`
	Exponent int     `xml:"exponent1,attr"`
}

type PolyCoef2DType struct {
	Coef      float64 `xml:",chardata"`
	Exponent1 int     `xml:"exponent1,attr"`
	Exponent2 int     `xml:"exponent2,attr"`
}

type XYZPolyType struct {
	X Poly1DType `xml:"X"`
	Y Poly1DType `xml:"Y"`
	Z Poly1DType `xml:"Z"`
}

type XYZPolyAttributeType struct {
	XYZPolyType
	Index int `xml:"index,attr"`
}

type Parameter struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

// Define types for enumerations
type Polarization1Type string
type Polarization2Type string
type DualPolarizationType string
type OrientationType string
// ImageFormation represents the ImageFormation element
type ImageFormation struct {
	RcvChanProc            RcvChanProcType           `xml:"RcvChanProc"`
	TxRcvPolarizationProc  DualPolarizationType      `xml:"TxRcvPolarizationProc"`
	TStartProc             float64                   `xml:"TStartProc"`
	TEndProc               float64                   `xml:"TEndProc"`
	TxFrequencyProc        TxFrequencyProcType       `xml:"TxFrequencyProc"`
	SegmentIdentifier      *string                   `xml:"SegmentIdentifier,omitempty"`
	ImageFormAlgo          string                    `xml:"ImageFormAlgo"`
	STBeamComp             string                    `xml:"STBeamComp"`
	ImageBeamComp          string                    `xml:"ImageBeamComp"`
	AzAutofocus            string                    `xml:"AzAutofocus"`
	RgAutofocus            string                    `xml:"RgAutofocus"`
	Processing             []ProcessingType          `xml:"Processing,omitempty"`
	PolarizationCalibration *PolarizationCalibrationType `xml:"PolarizationCalibration,omitempty"`
}

type RcvChanProcType struct {
	NumChanProc    int     `xml:"NumChanProc"`
	PRFScaleFactor *float64 `xml:"PRFScaleFactor,omitempty"`
	ChanIndex      []int   `xml:"ChanIndex"`
}

type TxFrequencyProcType struct {
	MinProc float64 `xml:"MinProc"`
	MaxProc float64 `xml:"MaxProc"`
}

type ProcessingType struct {
	Type     string      `xml:"Type"`
	Applied  bool        `xml:"Applied"`
	Parameter []Parameter `xml:"Parameter,omitempty"`
}

type PolarizationCalibrationType struct {
	DistortCorrectionApplied bool        `xml:"DistortCorrectionApplied"`
	Distortion               DistortionType `xml:"Distortion"`
}

type DistortionType struct {
	CalibrationDate *time.Time `xml:"CalibrationDate,omitempty"`
	A               float64    `xml:"A"`
	F1              ComplexType `xml:"F1"`
	Q1              ComplexType `xml:"Q1"`
	Q2              ComplexType `xml:"Q2"`
	F2              ComplexType `xml:"F2"`
	Q3              ComplexType `xml:"Q3"`
	Q4              ComplexType `xml:"Q4"`
	GainErrorA      *float64   `xml:"GainErrorA,omitempty"`
	GainErrorF1     *float64   `xml:"GainErrorF1,omitempty"`
	GainErrorF2     *float64   `xml:"GainErrorF2,omitempty"`
	PhaseErrorF1    *float64   `xml:"PhaseErrorF1,omitempty"`
	PhaseErrorF2    *float64   `xml:"PhaseErrorF2,omitempty"`
}

// SCPCOA represents the SCPCOA element
type SCPCOA struct {
	SCPTime     float64 `xml:"SCPTime"`
	ARPPos      XYZType `xml:"ARPPos"`
	ARPVel      XYZType `xml:"ARPVel"`
	ARPAcc      XYZType `xml:"ARPAcc"`
	SideOfTrack string  `xml:"SideOfTrack"`
	SlantRange  float64 `xml:"SlantRange"`
	GroundRange float64 `xml:"GroundRange"`
	DopplerConeAng float64 `xml:"DopplerConeAng"`
	GrazeAng    float64 `xml:"GrazeAng"`
	IncidenceAng float64 `xml:"IncidenceAng"`
	TwistAng    float64 `xml:"TwistAng"`
	SlopeAng    float64 `xml:"SlopeAng"`
	AzimAng     float64 `xml:"AzimAng"`
	LayoverAng  float64 `xml:"LayoverAng"`
	Bistatic    *BistaticType `xml:"Bistatic,omitempty"`
}

type BistaticType struct {
	BistaticAng     float64 `xml:"BistaticAng"`
	BistaticAngRate float64 `xml:"BistaticAngRate"`
	TxPlatform      SCPCOABistaticPlatformType `xml:"TxPlatform"`
	RcvPlatform     SCPCOABistaticPlatformType `xml:"RcvPlatform"`
}

type SCPCOABistaticPlatformType struct {
	Time      float64 `xml:"Time"`
	Pos       XYZType `xml:"Pos"`
	Vel       XYZType `xml:"Vel"`
	Acc       XYZType `xml:"Acc"`
	AzimAng   float64 `xml:"AzimAng"`
}

// Radiometric represents the Radiometric element
type Radiometric struct {
	NoiseLevel    *NoiseLevelType `xml:"NoiseLevel,omitempty"`
	RCSSFPoly     *Poly2DType     `xml:"RCSSFPoly,omitempty"`
	SigmaZeroSFPoly *Poly2DType    `xml:"SigmaZeroSFPoly,omitempty"`
	BetaZeroSFPoly *Poly2DType    `xml:"BetaZeroSFPoly,omitempty"`
	GammaZeroSFPoly *Poly2DType   `xml:"GammaZeroSFPoly,omitempty"`
}

type NoiseLevelType struct {
	NoiseLevelType string     `xml:"NoiseLevelType"`
	NoisePoly      Poly2DType `xml:"NoisePoly"`
}

// Antenna represents the Antenna element
type Antenna struct {
	Tx     *AntParamType `xml:"Tx,omitempty"`
	Rcv    *AntParamType `xml:"Rcv,omitempty"`
	TwoWay *AntParamType `xml:"TwoWay,omitempty"`
}

type AntParamType struct {
	XAxisPoly    XYZPolyType      `xml:"XAxisPoly"`
	YAxisPoly    XYZPolyType      `xml:"YAxisPoly"`
	FreqZero     float64          `xml:"FreqZero"`
	EB           *EBType          `xml:"EB,omitempty"`
	Array        GainPhasePolyType `xml:"Array"`
	Elem         *GainPhasePolyType `xml:"Elem,omitempty"`
	GainBSPoly   *Poly1DType      `xml:"GainBSPoly,omitempty"`
	EBFreqShift  *bool            `xml:"EBFreqShift,omitempty"`
	MLFreqDilation *bool          `xml:"MLFreqDilation,omitempty"`
}

type EBType struct {
	DCXPoly Poly1DType `xml:"DCXPoly"`
	DCYPoly Poly1DType `xml:"DCYPoly"`
}

type GainPhasePolyType struct {
	GainPoly  Poly2DType `xml:"GainPoly"`
	PhasePoly Poly2DType `xml:"PhasePoly"`
}

// ErrorStatistics represents the ErrorStatistics element
type ErrorStatistics struct {
	CompositeSCP        *CompositeErrorType       `xml:"CompositeSCP,omitempty"`
	Components          *ComponentsErrorType      `xml:"Components,omitempty"`
	BistaticCompositeSCP *BistaticCompositeErrorType `xml:"BistaticCompositeSCP,omitempty"`
	BistaticComponents   *BistaticComponentsErrorType `xml:"BistaticComponents,omitempty"`
	Unmodeled           *UnmodeledType            `xml:"Unmodeled,omitempty"`
	AdditionalParms     *AdditionalParmsType      `xml:"AdditionalParms,omitempty"`
	AdjustableParameterOffsets *AdjustableParameterOffsetsType `xml:"AdjustableParameterOffsets,omitempty"`
	BistaticAdjustableParameterOffsets *BistaticAdjustableParameterOffsetsType `xml:"BistaticAdjustableParameterOffsets,omitempty"`
}

// Define the remaining types (CompositeErrorType, ComponentsErrorType, etc.) as needed

// MatchInfo represents the MatchInfo element
type MatchInfo struct {
	NumMatchTypes int           `xml:"NumMatchTypes"`
	MatchType     []MatchTypeType `xml:"MatchType"`
}

type MatchTypeType struct {
	TypeID              string             `xml:"TypeID"`
	CurrentIndex        *int               `xml:"CurrentIndex,omitempty"`
	NumMatchCollections int                `xml:"NumMatchCollections"`
	MatchCollection     []MatchCollectionType `xml:"MatchCollection,omitempty"`
	Index               int                `xml:"index,attr"`
}

type MatchCollectionType struct {
	CoreName   string      `xml:"CoreName"`
	MatchIndex *int        `xml:"MatchIndex,omitempty"`
	Parameter  []Parameter `xml:"Parameter,omitempty"`
	Index      int         `xml:"index,attr"`
}

// RgAzComp represents the RgAzComp element
type RgAzComp struct {
	AzSF    float64    `xml:"AzSF"`
	KazPoly Poly1DType `xml:"KazPoly"`
}

// PFA represents the PFA element
type PFA struct {
	FPN              XYZType    `xml:"FPN"`
	IPN              XYZType    `xml:"IPN"`
	PolarAngRefTime  float64    `xml:"PolarAngRefTime"`
	PolarAngPoly     Poly1DType `xml:"PolarAngPoly"`
	SpatialFreqSFPoly Poly1DType `xml:"SpatialFreqSFPoly"`
	Krg1             float64    `xml:"Krg1"`
	Krg2             float64    `xml:"Krg2"`
	Kaz1             float64    `xml:"Kaz1"`
	Kaz2             float64    `xml:"Kaz2"`
	STDeskew         *STDeskewType `xml:"STDeskew,omitempty"`
}

type STDeskewType struct {
	Applied       bool       `xml:"Applied"`
	STDSPhasePoly Poly2DType `xml:"STDSPhasePoly"`
}

// RMA represents the RMA element
type RMA struct {
	RMAlgoType string    `xml:"RMAlgoType"`
	ImageType  string    `xml:"ImageType"`
	RMAT       *RMATType `xml:"RMAT,omitempty"`
	RMCR       *RMCRType `xml:"RMCR,omitempty"`
	INCA       *INCAType `xml:"INCA,omitempty"`
}

type RMATType struct {
	PosRef         XYZType `xml:"PosRef"`
	VelRef         XYZType `xml:"VelRef"`
	DopConeAngRef  float64 `xml:"DopConeAngRef"`
}

type RMCRType struct {
	PosRef         XYZType `xml:"PosRef"`
	VelRef         XYZType `xml:"VelRef"`
	DopConeAngRef  float64 `xml:"DopConeAngRef"`
}

type INCAType struct {
	TimeCAPoly     Poly1DType `xml:"TimeCAPoly"`
	R_CA_SCP       float64    `xml:"R_CA_SCP"`
	FreqZero       float64    `xml:"FreqZero"`
	DRateSFPoly    Poly2DType `xml:"DRateSFPoly"`
	DopCentroidPoly *Poly2DType `xml:"DopCentroidPoly,omitempty"`
	DopCentroidCOA *bool      `xml:"DopCentroidCOA,omitempty"`
}

// Additional types used above
type ComplexType struct {
	Real float64 `xml:"Real"`
	Imag float64 `xml:"Imag"`
}

type AdditionalParmsType struct {
	Parameter []Parameter `xml:"Parameter"`
}

type AdjustableParameterOffsetsType struct {
	ARPPosSCPCOA   XYZType           `xml:"ARPPosSCPCOA"`
	ARPVel         XYZType           `xml:"ARPVel"`
	TxTimeSCPCOA   float64           `xml:"TxTimeSCPCOA"`
	RcvTimeSCPCOA  float64           `xml:"RcvTimeSCPCOA"`
	APOError       *Matrix2DType     `xml:"APOError,omitempty"`
	CompositeSCP   *CompositeErrorType `xml:"CompositeSCP,omitempty"`
}

type BistaticAdjustableParameterOffsetsType struct {
	TxPlatform     BistaticAdjustableParametersType `xml:"TxPlatform"`
	RcvPlatform    BistaticAdjustableParametersType `xml:"RcvPlatform"`
	APOError       *Matrix2DType                    `xml:"APOError,omitempty"`
	BistaticCompositeSCP *BistaticCompositeErrorType `xml:"BistaticCompositeSCP,omitempty"`
}

type BistaticAdjustableParametersType struct {
	APCPosSCPCOA XYZType `xml:"APCPosSCPCOA"`
	APCVel       XYZType `xml:"APCVel"`
	TimeSCPCOA   float64 `xml:"TimeSCPCOA"`
	ClockFreqSF  float64 `xml:"ClockFreqSF"`
}

type Matrix2DType struct {
	Entry []MatrixEntry2DType `xml:"Entry"`
}

type MatrixEntry2DType struct {
	Value  float64 `xml:",chardata"`
	Index1 int     `xml:"index1,attr"`
	Index2 int     `xml:"index2,attr"`
}
