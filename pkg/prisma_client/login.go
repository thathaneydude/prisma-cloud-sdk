package prisma_client

//import (
//	"PrismaCloud/pkg"
//	"PrismaCloud/pkg/constants"
//	"PrismaCloud/pkg/cwpp"
//	"fmt"
//)
//
//func (c *PrismaCloudClient) Login(username string, password string, customerName string, prismaId string) error {
//	loginResponse, err := c.Cspm.Login(username, password, customerName, prismaId)
//	if err != nil {
//		return err
//	}
//
//	c.Cspm.SetHeader(constants.AuthHeader, loginResponse.Token)
//	err = c.initializeCwpp(c.cwppApiVersion)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (c *PrismaCloudClient) initializeCwpp(apiVersion string) error {
//	resp, err := c.GetMetaInfo()
//	if err != nil {
//		return err
//	}
//	c.cwppBaseUrl = resp.TwistlockUrl
//	c.Cwpp, err = cwpp.NewCwppClient(c.cwppBaseUrl, apiVersion, c.sslVerify, c.schema)
//	if err != nil {
//		return &pkg.GenericError{Msg: fmt.Sprintf("Failed to initialize CWPP client using meta_info: %v", err)}
//	}
//	return nil
//}
