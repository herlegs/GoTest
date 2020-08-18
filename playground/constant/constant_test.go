package constant

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
	"testing"
)

var cf = `ADS
DATA_ENRICHMENT
APEX
APEX_AUTOPLAT-INFRA
APEX-AUTOPLAT-DEPLOY
APEX-AUTOPLAT-INFRA
APEX-AUTOPLAT-TEST
APEX-CF
APEX-NETVIZ
APEX-NETVIZ-OBS
APEX-COSMOS
COSMOS
CLUSTERMANAGEMENT
DBOPS
DBOps
DEVEXP
devexp
lloyd
lloyd2
NIST
OBSERVABILITY
OPSTOOL
QA
SRE
SRE-CORE
SYSOPS
Sysops
TECHTRAINING
terratest
workshop
APEX-NEXUS
Bali
BALI
CAULDRON
Cauldron
cauldron
data-science
DATASCIENCE
SERVICESANDDATA
CX
CSE
data-engineering
DATAENG
DATA_INSIGHTS
SEGMENTATION
DS Food
FOOD
GRABEXPRESS
GRABFOOD
LOGISTICS
DAX
DAXEXP
ENTERPRISE
FDW
SNT
SnT
GEO
geo
VENTURES
venture
WHEELS
DOTCOM
GRABKIOS
GRABPLATFORM
IT
IT_Voice
LOYALTY
REWARDS
MARKETING
ECONS
MARKETPLACE
Marketplace
marketplace
MEX
MPT
BPA
GRABPAY
PAYMENTS
Payments
POPSTOOLS
EXPERIMENTATION
Experimentation
RTC
SENTRY
SAPTF
GRAB-ID
GRABID
malacca
USERTRUST
USERTURST
USERTRUST_SHARED
SAFETY
SECURITY
security
GRABHITCH
GRABNOW
GRABSHARE
GRABSHUTTLE
TRANSPORT
GROWTH
SHARED`

var tf = `Ads & Data Monetisation
Ads & Data Monetisation
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
APEX
Bali
Bali
Cauldron
Cauldron
Cauldron
Cauldron
Cauldron
Coban
Consumer Experience (CX)
Customer Support Experience (CSE)
Data Engineering
Data Engineering
Data Insights
Data Insights
Deliveries
Deliveries
Deliveries
Deliveries
Deliveries
Driver Experience (Dax Exp)
Driver Experience (Dax Exp)
Enterprise
Finance Cortex
Gamma/SNT
Gamma/SNT
Geo
Geo
Grab Ventures
Grab Wheels
Grab Wheels
Grab.com
GrabKiosk
GrabPlatform
IT
IT
Loyalty
Loyalty
Marketing
Marketplace
Marketplace
Marketplace
Marketplace
Merchant Experience
Mobile Platform
OPERATIONS
Payment Experience
Payment Experience
Payment Experience
People Tech
Product Insights & Experimentation (PIE)
Product Insights & Experimentation (PIE)
Real Time Communication (RTC)
Sentry
SuperApp Partnerships TF
TIS
TIS
TIS
TIS
TIS
TIS
TISS
TISS
TISS
Transport
Transport
Transport
Transport
Transport
Unowned
Unowned`


var tfMap =`ADS          = ADS & DATA MONETISATION
APEX         = APEX
BALI         = BALI
CAULDRON     = CAULDRON
CSE          = CUSTOMER SUPPORT EXPERIENCE (CSE)
CX           = CONSUMER EXPERIENCE (CX)
DE           = DATA ENGINEERING
DAXEXP       = DRIVER EXPERIENCE (DAX EXP)
DELIVERIES   = DELIVERIES
MARKETPLACE  = MARKETPLACE
ENTERPRISE   = ENTERPRISE
PIE          = PRODUCT INSIGHTS & EXPERIMENTATION (PIE)
GEO          = GEO
TRANSPORT    = TRANSPORT
TIS          = TIS
PE           = PAYMENT EXPERIENCE
GRABPLATFORM = GRABPLATFORM
UNOWNED      = UNOWNED
IT           = IT
LOYALTY      = LOYALTY
MARKETING    = MARKETING
MP           = MOBILE PLATFORM
PT           = PEOPLE TECH
RTC          = REAL TIME COMMUNICATION (RTC)
SP           = SUPERAPP PARTNERSHIPS TF
SECURITY     = SECURITY
DI           = DATA INSIGHTS
COBAN        = COBAN
GAMMA        = GAMMA SNT
GW           = GRAB WHEELS
GRABCOM      = GRAB.COM
AWS          = AI/ML INFRA (AWS)
FC           = FINANCE CORTEX
OPERATIONS   = OPERATIONS
PD           = PRODUCT DESIGN`

var tfNameMap = map[string]string{}
var once sync.Once

// GetTFVarName
func GetTFVarName(tf string) string {
	once.Do(func() {
		for _, l := range strings.Split(tfMap,"\n") {
			parts := strings.Split(l, "=")
			for i, p := range parts {
				parts[i] = strings.TrimSpace(p)
			}
			name, val := parts[0], parts[1]
			tfNameMap[val] = name
		}
	})
	return tfNameMap[tf]
}

func TestName(t *testing.T) {
	cfs := strings.Split(cf, "\n")
	tfs := strings.Split(tf, "\n")
	assert.Equal(t,len(cfs), len(tfs))
	consts := ""
	maps := ""
	dedup := map[string]bool{}
	for i,p := range cfs {
		cfVarName := normal(p)
		if dedup[cfVarName] {
			continue
		}
		dedup[cfVarName] = true
		tfVal := strings.ToUpper(tfs[i])
		if !TechFamilies[tfVal] {
			fmt.Printf("!!!%v, cloest one: %v\n",cfVarName+" :"+tfVal+" is not right",findClosestTF(tfVal))
		}
		consts += fmt.Sprintf("CF_%v = \"%v\"\n",cfVarName,strings.ToUpper(p))
		maps += fmt.Sprintf("CF_%v:%v,\n",cfVarName, GetTFVarName(tfVal))
	}
	fmt.Printf("%v\n",maps)
}

// normal  
func normal(str string) string {
	str = strings.ReplaceAll(str, " ", "_")
	str = strings.ReplaceAll(str, "-", "_")
	return strings.ToUpper(str)
}


// longestCommonSubsequence ...
func longestCommonSubsequence(a, b string) int {
	a, b = strings.ToUpper(a), strings.ToUpper(b)
	m, n := len(a), len(b)
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i%2][j] = dp[(i-1)%2][j-1] + 1
			}
			dp[i%2][j] = max(dp[i%2][j], dp[(i-1)%2][j], dp[i%2][j-1])
		}
	}
	return dp[m%2][n]
}

func max(a int, list ...int) int {
	for _, i := range list {
		if i > a {
			a = i
		}
	}
	return a
}

// findClosestTF ...
func findClosestTF(tf string) string {
	lcs := 0
	closest := ""
	for techFamily := range TechFamilies {
		if l := longestCommonSubsequence(tf, techFamily); l > lcs {
			lcs = l
			closest = techFamily
		}
	}
	return closest
}