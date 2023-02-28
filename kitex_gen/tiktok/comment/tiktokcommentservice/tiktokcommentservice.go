// Code generated by Kitex v0.4.4. DO NOT EDIT.

package tiktokcommentservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	comment "github.com/ozline/tiktok/kitex_gen/tiktok/comment"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return tiktokCommentServiceServiceInfo
}

var tiktokCommentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "tiktokCommentService"
	handlerType := (*comment.TiktokCommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"List":         kitex.NewMethodInfo(listHandler, newListArgs, newListResult, false),
		"Post":         kitex.NewMethodInfo(postHandler, newPostArgs, newPostResult, false),
		"SetLike":      kitex.NewMethodInfo(setLikeHandler, newSetLikeArgs, newSetLikeResult, false),
		"GetLike":      kitex.NewMethodInfo(getLikeHandler, newGetLikeArgs, newGetLikeResult, false),
		"SetFavorite":  kitex.NewMethodInfo(setFavoriteHandler, newSetFavoriteArgs, newSetFavoriteResult, false),
		"FavoriteList": kitex.NewMethodInfo(favoriteListHandler, newFavoriteListArgs, newFavoriteListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func listHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.ListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.TiktokCommentService).List(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ListArgs:
		success, err := handler.(comment.TiktokCommentService).List(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListResult)
		realResult.Success = success
	}
	return nil
}
func newListArgs() interface{} {
	return &ListArgs{}
}

func newListResult() interface{} {
	return &ListResult{}
}

type ListArgs struct {
	Req *comment.ListReq
}

func (p *ListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.ListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ListArgs) Unmarshal(in []byte) error {
	msg := new(comment.ListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListArgs_Req_DEFAULT *comment.ListReq

func (p *ListArgs) GetReq() *comment.ListReq {
	if !p.IsSetReq() {
		return ListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListArgs) IsSetReq() bool {
	return p.Req != nil
}

type ListResult struct {
	Success *comment.ListResp
}

var ListResult_Success_DEFAULT *comment.ListResp

func (p *ListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.ListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ListResult) Unmarshal(in []byte) error {
	msg := new(comment.ListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListResult) GetSuccess() *comment.ListResp {
	if !p.IsSetSuccess() {
		return ListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.ListResp)
}

func (p *ListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func postHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.PostReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.TiktokCommentService).Post(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PostArgs:
		success, err := handler.(comment.TiktokCommentService).Post(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PostResult)
		realResult.Success = success
	}
	return nil
}
func newPostArgs() interface{} {
	return &PostArgs{}
}

func newPostResult() interface{} {
	return &PostResult{}
}

type PostArgs struct {
	Req *comment.PostReq
}

func (p *PostArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.PostReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PostArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PostArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PostArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PostArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PostArgs) Unmarshal(in []byte) error {
	msg := new(comment.PostReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PostArgs_Req_DEFAULT *comment.PostReq

func (p *PostArgs) GetReq() *comment.PostReq {
	if !p.IsSetReq() {
		return PostArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PostArgs) IsSetReq() bool {
	return p.Req != nil
}

type PostResult struct {
	Success *comment.PostResp
}

var PostResult_Success_DEFAULT *comment.PostResp

func (p *PostResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.PostResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PostResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PostResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PostResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PostResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PostResult) Unmarshal(in []byte) error {
	msg := new(comment.PostResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PostResult) GetSuccess() *comment.PostResp {
	if !p.IsSetSuccess() {
		return PostResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PostResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.PostResp)
}

func (p *PostResult) IsSetSuccess() bool {
	return p.Success != nil
}

func setLikeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.LikeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.TiktokCommentService).SetLike(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SetLikeArgs:
		success, err := handler.(comment.TiktokCommentService).SetLike(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SetLikeResult)
		realResult.Success = success
	}
	return nil
}
func newSetLikeArgs() interface{} {
	return &SetLikeArgs{}
}

func newSetLikeResult() interface{} {
	return &SetLikeResult{}
}

type SetLikeArgs struct {
	Req *comment.LikeReq
}

func (p *SetLikeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.LikeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SetLikeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SetLikeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SetLikeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetLikeArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SetLikeArgs) Unmarshal(in []byte) error {
	msg := new(comment.LikeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SetLikeArgs_Req_DEFAULT *comment.LikeReq

func (p *SetLikeArgs) GetReq() *comment.LikeReq {
	if !p.IsSetReq() {
		return SetLikeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetLikeArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetLikeResult struct {
	Success *comment.LikeResp
}

var SetLikeResult_Success_DEFAULT *comment.LikeResp

func (p *SetLikeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.LikeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SetLikeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SetLikeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SetLikeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetLikeResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SetLikeResult) Unmarshal(in []byte) error {
	msg := new(comment.LikeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetLikeResult) GetSuccess() *comment.LikeResp {
	if !p.IsSetSuccess() {
		return SetLikeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetLikeResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.LikeResp)
}

func (p *SetLikeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getLikeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.LikeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.TiktokCommentService).GetLike(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetLikeArgs:
		success, err := handler.(comment.TiktokCommentService).GetLike(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetLikeResult)
		realResult.Success = success
	}
	return nil
}
func newGetLikeArgs() interface{} {
	return &GetLikeArgs{}
}

func newGetLikeResult() interface{} {
	return &GetLikeResult{}
}

type GetLikeArgs struct {
	Req *comment.LikeReq
}

func (p *GetLikeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.LikeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetLikeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetLikeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetLikeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetLikeArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetLikeArgs) Unmarshal(in []byte) error {
	msg := new(comment.LikeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetLikeArgs_Req_DEFAULT *comment.LikeReq

func (p *GetLikeArgs) GetReq() *comment.LikeReq {
	if !p.IsSetReq() {
		return GetLikeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLikeArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetLikeResult struct {
	Success *comment.LikeResp
}

var GetLikeResult_Success_DEFAULT *comment.LikeResp

func (p *GetLikeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.LikeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetLikeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetLikeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetLikeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetLikeResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetLikeResult) Unmarshal(in []byte) error {
	msg := new(comment.LikeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLikeResult) GetSuccess() *comment.LikeResp {
	if !p.IsSetSuccess() {
		return GetLikeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLikeResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.LikeResp)
}

func (p *GetLikeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func setFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.FavoriteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.TiktokCommentService).SetFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SetFavoriteArgs:
		success, err := handler.(comment.TiktokCommentService).SetFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SetFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newSetFavoriteArgs() interface{} {
	return &SetFavoriteArgs{}
}

func newSetFavoriteResult() interface{} {
	return &SetFavoriteResult{}
}

type SetFavoriteArgs struct {
	Req *comment.FavoriteReq
}

func (p *SetFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.FavoriteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SetFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SetFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SetFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SetFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(comment.FavoriteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SetFavoriteArgs_Req_DEFAULT *comment.FavoriteReq

func (p *SetFavoriteArgs) GetReq() *comment.FavoriteReq {
	if !p.IsSetReq() {
		return SetFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetFavoriteResult struct {
	Success *comment.FavoriteResp
}

var SetFavoriteResult_Success_DEFAULT *comment.FavoriteResp

func (p *SetFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.FavoriteResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SetFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SetFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SetFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SetFavoriteResult) Unmarshal(in []byte) error {
	msg := new(comment.FavoriteResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetFavoriteResult) GetSuccess() *comment.FavoriteResp {
	if !p.IsSetSuccess() {
		return SetFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.FavoriteResp)
}

func (p *SetFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.FavoriteListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.TiktokCommentService).FavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteListArgs:
		success, err := handler.(comment.TiktokCommentService).FavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteListArgs() interface{} {
	return &FavoriteListArgs{}
}

func newFavoriteListResult() interface{} {
	return &FavoriteListResult{}
}

type FavoriteListArgs struct {
	Req *comment.FavoriteListReq
}

func (p *FavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.FavoriteListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(comment.FavoriteListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteListArgs_Req_DEFAULT *comment.FavoriteListReq

func (p *FavoriteListArgs) GetReq() *comment.FavoriteListReq {
	if !p.IsSetReq() {
		return FavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoriteListResult struct {
	Success *comment.FavoriteListResp
}

var FavoriteListResult_Success_DEFAULT *comment.FavoriteListResp

func (p *FavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.FavoriteListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteListResult) Unmarshal(in []byte) error {
	msg := new(comment.FavoriteListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteListResult) GetSuccess() *comment.FavoriteListResp {
	if !p.IsSetSuccess() {
		return FavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.FavoriteListResp)
}

func (p *FavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) List(ctx context.Context, Req *comment.ListReq) (r *comment.ListResp, err error) {
	var _args ListArgs
	_args.Req = Req
	var _result ListResult
	if err = p.c.Call(ctx, "List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Post(ctx context.Context, Req *comment.PostReq) (r *comment.PostResp, err error) {
	var _args PostArgs
	_args.Req = Req
	var _result PostResult
	if err = p.c.Call(ctx, "Post", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetLike(ctx context.Context, Req *comment.LikeReq) (r *comment.LikeResp, err error) {
	var _args SetLikeArgs
	_args.Req = Req
	var _result SetLikeResult
	if err = p.c.Call(ctx, "SetLike", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetLike(ctx context.Context, Req *comment.LikeReq) (r *comment.LikeResp, err error) {
	var _args GetLikeArgs
	_args.Req = Req
	var _result GetLikeResult
	if err = p.c.Call(ctx, "GetLike", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetFavorite(ctx context.Context, Req *comment.FavoriteReq) (r *comment.FavoriteResp, err error) {
	var _args SetFavoriteArgs
	_args.Req = Req
	var _result SetFavoriteResult
	if err = p.c.Call(ctx, "SetFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, Req *comment.FavoriteListReq) (r *comment.FavoriteListResp, err error) {
	var _args FavoriteListArgs
	_args.Req = Req
	var _result FavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
