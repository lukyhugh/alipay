package alipay

import "context"

// FaceCertifyInitialize Web人脸核身初始化 https://opendocs.alipay.com/open/02zloa?scene=common&pathHash=b0b7fece
func (c *Client) FaceCertifyInitialize(ctx context.Context, param FaceCertifyInitialize) (result *FaceCertifyInitializeRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// FaceCertifyVerify Web人脸核身开始认证 https://opendocs.alipay.com/open/02zlob?scene=common&pathHash=20eba12a
func (c *Client) FaceCertifyVerify(ctx context.Context, param FaceCertifyVerify) (result *FaceCertifyVerifyRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// FaceCertifyQuery Web人脸核身记录查询 https://opendocs.alipay.com/open/02zloc?scene=common&pathHash=b1141506
func (c *Client) FaceCertifyQuery(ctx context.Context, param FaceCertifyQuery) (result *FaceCertifyQueryRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}
