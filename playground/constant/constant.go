package constant

// Tech Families
const (
	ADS          = "ADS & DATA MONETISATION"
	APEX         = "APEX"
	BALI         = "BALI"
	CAULDRON     = "CAULDRON"
	CSE          = "CUSTOMER SUPPORT EXPERIENCE (CSE)"
	CX           = "CONSUMER EXPERIENCE (CX)"
	DE           = "DATA ENGINEERING"
	DAXEXP       = "DRIVER EXPERIENCE (DAX EXP)"
	DELIVERIES   = "DELIVERIES"
	MARKETPLACE  = "MARKETPLACE"
	ENTERPRISE   = "ENTERPRISE"
	PIE          = "PRODUCT INSIGHTS & EXPERIMENTATION (PIE)"
	GEO          = "GEO"
	TRANSPORT    = "TRANSPORT"
	TIS          = "TIS"
	PE           = "PAYMENT EXPERIENCE"
	GRABPLATFORM = "GRABPLATFORM"
	UNOWNED      = "UNOWNED"
	IT           = "IT"
	LOYALTY      = "LOYALTY"
	MARKETING    = "MARKETING"
	MP           = "MOBILE PLATFORM"
	PT           = "PEOPLE TECH"
	RTC          = "REAL TIME COMMUNICATION (RTC)"
	SP           = "SUPERAPP PARTNERSHIPS TF"
	SECURITY     = "SECURITY"
	DI           = "DATA INSIGHTS"
	COBAN        = "COBAN"
	GAMMA        = "GAMMA SNT"
	GW           = "GRAB WHEELS"
	GRABCOM      = "GRAB.COM"
	AWS          = "AI/ML INFRA (AWS)"
	FC           = "FINANCE CORTEX"
	OPERATIONS   = "OPERATIONS"
	PD           = "PRODUCT DESIGN"
)

var TechFamilies = map[string]bool{
	ADS:          true,
	APEX:         true,
	BALI:         true,
	CAULDRON:     true,
	CSE:          true,
	CX:           true,
	DE:           true,
	DAXEXP:       true,
	DELIVERIES:   true,
	MARKETPLACE:  true,
	ENTERPRISE:   true,
	PIE:          true,
	GEO:          true,
	TRANSPORT:    true,
	TIS:          true,
	PE:           true,
	GRABPLATFORM: true,
	UNOWNED:      true,
	IT:           true,
	LOYALTY:      true,
	MARKETING:    true,
	MP:           true,
	PT:           true,
	RTC:          true,
	SP:           true,
	SECURITY:     true,
	DI:           true,
	COBAN:        true,
	GAMMA:        true,
	GW:           true,
	GRABCOM:      true,
	AWS:          true,
	FC:           true,
	OPERATIONS:   true,
	PD:           true,
}

// Costing Families
const (
	CF_ADS                  = "ADS"
	CF_DATA_ENRICHMENT      = "DATA_ENRICHMENT"
	CF_APEX                 = "APEX"
	CF_APEX_AUTOPLAT_INFRA  = "APEX_AUTOPLAT-INFRA"
	CF_APEX_AUTOPLAT_DEPLOY = "APEX-AUTOPLAT-DEPLOY"
	CF_APEX_AUTOPLAT_TEST   = "APEX-AUTOPLAT-TEST"
	CF_APEX_CF              = "APEX-CF"
	CF_APEX_NETVIZ          = "APEX-NETVIZ"
	CF_APEX_NETVIZ_OBS      = "APEX-NETVIZ-OBS"
	CF_APEX_COSMOS          = "APEX-COSMOS"
	CF_COSMOS               = "COSMOS"
	CF_CLUSTERMANAGEMENT    = "CLUSTERMANAGEMENT"
	CF_DBOPS                = "DBOPS"
	CF_DEVEXP               = "DEVEXP"
	CF_LLOYD                = "LLOYD"
	CF_LLOYD2               = "LLOYD2"
	CF_NIST                 = "NIST"
	CF_OBSERVABILITY        = "OBSERVABILITY"
	CF_OPSTOOL              = "OPSTOOL"
	CF_QA                   = "QA"
	CF_SRE                  = "SRE"
	CF_SRE_CORE             = "SRE-CORE"
	CF_SYSOPS               = "SYSOPS"
	CF_TECHTRAINING         = "TECHTRAINING"
	CF_TERRATEST            = "TERRATEST"
	CF_WORKSHOP             = "WORKSHOP"
	CF_APEX_NEXUS           = "APEX-NEXUS"
	CF_BALI                 = "BALI"
	CF_CAULDRON             = "CAULDRON"
	CF_DATA_SCIENCE         = "DATA-SCIENCE"
	CF_DATASCIENCE          = "DATASCIENCE"
	CF_SERVICESANDDATA      = "SERVICESANDDATA"
	CF_CX                   = "CX"
	CF_CSE                  = "CSE"
	CF_DATA_ENGINEERING     = "DATA-ENGINEERING"
	CF_DATAENG              = "DATAENG"
	CF_DATA_INSIGHTS        = "DATA_INSIGHTS"
	CF_SEGMENTATION         = "SEGMENTATION"
	CF_DS_FOOD              = "DS FOOD"
	CF_FOOD                 = "FOOD"
	CF_GRABEXPRESS          = "GRABEXPRESS"
	CF_GRABFOOD             = "GRABFOOD"
	CF_LOGISTICS            = "LOGISTICS"
	CF_DAX                  = "DAX"
	CF_DAXEXP               = "DAXEXP"
	CF_ENTERPRISE           = "ENTERPRISE"
	CF_FDW                  = "FDW"
	CF_SNT                  = "SNT"
	CF_GEO                  = "GEO"
	CF_VENTURES             = "VENTURES"
	CF_VENTURE              = "VENTURE"
	CF_WHEELS               = "WHEELS"
	CF_DOTCOM               = "DOTCOM"
	CF_GRABKIOS             = "GRABKIOS"
	CF_GRABPLATFORM         = "GRABPLATFORM"
	CF_IT                   = "IT"
	CF_IT_VOICE             = "IT_VOICE"
	CF_LOYALTY              = "LOYALTY"
	CF_REWARDS              = "REWARDS"
	CF_MARKETING            = "MARKETING"
	CF_ECONS                = "ECONS"
	CF_MARKETPLACE          = "MARKETPLACE"
	CF_MEX                  = "MEX"
	CF_MPT                  = "MPT"
	CF_BPA                  = "BPA"
	CF_GRABPAY              = "GRABPAY"
	CF_PAYMENTS             = "PAYMENTS"
	CF_POPSTOOLS            = "POPSTOOLS"
	CF_EXPERIMENTATION      = "EXPERIMENTATION"
	CF_RTC                  = "RTC"
	CF_SENTRY               = "SENTRY"
	CF_SAPTF                = "SAPTF"
	CF_GRAB_ID              = "GRAB-ID"
	CF_GRABID               = "GRABID"
	CF_MALACCA              = "MALACCA"
	CF_USERTRUST            = "USERTRUST"
	CF_USERTURST            = "USERTURST"
	CF_USERTRUST_SHARED     = "USERTRUST_SHARED"
	CF_SAFETY               = "SAFETY"
	CF_SECURITY             = "SECURITY"
	CF_GRABHITCH            = "GRABHITCH"
	CF_GRABNOW              = "GRABNOW"
	CF_GRABSHARE            = "GRABSHARE"
	CF_GRABSHUTTLE          = "GRABSHUTTLE"
	CF_TRANSPORT            = "TRANSPORT"
	CF_GROWTH               = "GROWTH"
	CF_SHARED               = "SHARED"
)

var CostingFamilyTFMap = map[string]string{
	CF_ADS:                  ADS,
	CF_DATA_ENRICHMENT:      ADS,
	CF_APEX:                 APEX,
	CF_APEX_AUTOPLAT_INFRA:  APEX,
	CF_APEX_AUTOPLAT_DEPLOY: APEX,
	CF_APEX_AUTOPLAT_TEST:   APEX,
	CF_APEX_CF:              APEX,
	CF_APEX_NETVIZ:          APEX,
	CF_APEX_NETVIZ_OBS:      APEX,
	CF_APEX_COSMOS:          APEX,
	CF_COSMOS:               APEX,
	CF_CLUSTERMANAGEMENT:    APEX,
	CF_DBOPS:                APEX,
	CF_DEVEXP:               APEX,
	CF_LLOYD:                APEX,
	CF_LLOYD2:               APEX,
	CF_NIST:                 APEX,
	CF_OBSERVABILITY:        APEX,
	CF_OPSTOOL:              APEX,
	CF_QA:                   APEX,
	CF_SRE:                  APEX,
	CF_SRE_CORE:             APEX,
	CF_SYSOPS:               APEX,
	CF_TECHTRAINING:         APEX,
	CF_TERRATEST:            APEX,
	CF_WORKSHOP:             APEX,
	CF_APEX_NEXUS:           APEX,
	CF_BALI:                 BALI,
	CF_CAULDRON:             CAULDRON,
	CF_DATA_SCIENCE:         CAULDRON,
	CF_DATASCIENCE:          CAULDRON,
	CF_SERVICESANDDATA:      COBAN,
	CF_CX:                   CX,
	CF_CSE:                  CSE,
	CF_DATA_ENGINEERING:     DE,
	CF_DATAENG:              DE,
	CF_DATA_INSIGHTS:        DI,
	CF_SEGMENTATION:         DI,
	CF_DS_FOOD:              DELIVERIES,
	CF_FOOD:                 DELIVERIES,
	CF_GRABEXPRESS:          DELIVERIES,
	CF_GRABFOOD:             DELIVERIES,
	CF_LOGISTICS:            DELIVERIES,
	CF_DAX:                  DAXEXP,
	CF_DAXEXP:               DAXEXP,
	CF_ENTERPRISE:           ENTERPRISE,
	CF_FDW:                  FC,
	CF_SNT:                  GAMMA,
	CF_GEO:                  GEO,
	CF_VENTURE:              GW,
	CF_WHEELS:               GW,
	CF_DOTCOM:               GRABCOM,
	CF_GRABPLATFORM:         GRABPLATFORM,
	CF_IT:                   IT,
	CF_IT_VOICE:             IT,
	CF_LOYALTY:              LOYALTY,
	CF_REWARDS:              LOYALTY,
	CF_MARKETING:            MARKETING,
	CF_ECONS:                MARKETPLACE,
	CF_MARKETPLACE:          MARKETPLACE,
	CF_MPT:                  MP,
	CF_BPA:                  OPERATIONS,
	CF_GRABPAY:              PE,
	CF_PAYMENTS:             PE,
	CF_POPSTOOLS:            PT,
	CF_EXPERIMENTATION:      PIE,
	CF_RTC:                  RTC,
	CF_SAPTF:                SP,
	CF_GRAB_ID:              TIS,
	CF_GRABID:               TIS,
	CF_MALACCA:              TIS,
	CF_USERTRUST:            TIS,
	CF_USERTURST:            TIS,
	CF_USERTRUST_SHARED:     TIS,
	CF_SAFETY:               TIS,
	CF_SECURITY:             TIS,
	CF_GRABHITCH:            TRANSPORT,
	CF_GRABNOW:              TRANSPORT,
	CF_GRABSHARE:            TRANSPORT,
	CF_GRABSHUTTLE:          TRANSPORT,
	CF_TRANSPORT:            TRANSPORT,
	CF_GROWTH:               UNOWNED,
	CF_SHARED:               UNOWNED,
}

var TFReverse = map[string]string{}
