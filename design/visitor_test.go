package design

import "testing"

func TestServiceRequestVisitor_Visit(t *testing.T) {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Add(NewEnterpriseCustomer("bob"))
	c.Add(NewIndividualCustomer("废弃工厂"))
	c.Accept(&ServiceRequestVisitor{})
}

func TestAnalysis(t *testing.T){
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Add(NewEnterpriseCustomer("bob"))
	c.Add(NewIndividualCustomer("废弃工厂"))
	c.Accept(&AnalysisVisitor{})
}