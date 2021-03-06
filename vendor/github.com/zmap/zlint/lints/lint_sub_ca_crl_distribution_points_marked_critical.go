package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

/************************************************
BRs: 7.1.2.2b cRLDistributionPoints
This extension MUST be present and MUST NOT be marked critical.
It MUST contain the HTTP URL of the CA’s CRL service.
************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCACRLDistCrit struct{}

func (l *subCACRLDistCrit) Initialize() error {
	return nil
}

func (l *subCACRLDistCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.CrlDistOID)
}

func (l *subCACRLDistCrit) Execute(c *x509.Certificate) *LintResult {
	if e := util.GetExtFromCert(c, util.CrlDistOID); e.Critical {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_crl_distribution_points_marked_critical",
		Description:   "Subordinate CA Certificate: cRLDistributionPoints MUST be present and MUST NOT be marked critical.",
		Citation:      "BRs: 7.1.2.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCACRLDistCrit{},
	})
}
