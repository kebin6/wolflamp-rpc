// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/kebin6/wolflamp-rpc/ent/banner"
	"github.com/kebin6/wolflamp-rpc/ent/exchange"
	"github.com/kebin6/wolflamp-rpc/ent/file"
	"github.com/kebin6/wolflamp-rpc/ent/order"
	"github.com/kebin6/wolflamp-rpc/ent/origininvitecode"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/ent/pool"
	"github.com/kebin6/wolflamp-rpc/ent/reward"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
	"github.com/kebin6/wolflamp-rpc/ent/setting"
	"github.com/kebin6/wolflamp-rpc/ent/statement"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type BannerPager struct {
	Order  banner.OrderOption
	Filter func(*BannerQuery) (*BannerQuery, error)
}

// BannerPaginateOption enables pagination customization.
type BannerPaginateOption func(*BannerPager)

// DefaultBannerOrder is the default ordering of Banner.
var DefaultBannerOrder = Desc(banner.FieldID)

func newBannerPager(opts []BannerPaginateOption) (*BannerPager, error) {
	pager := &BannerPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultBannerOrder
	}
	return pager, nil
}

func (p *BannerPager) ApplyFilter(query *BannerQuery) (*BannerQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// BannerPageList is Banner PageList result.
type BannerPageList struct {
	List        []*Banner    `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (b *BannerQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...BannerPaginateOption,
) (*BannerPageList, error) {

	pager, err := newBannerPager(opts)
	if err != nil {
		return nil, err
	}

	if b, err = pager.ApplyFilter(b); err != nil {
		return nil, err
	}

	ret := &BannerPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := b.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		b = b.Order(pager.Order)
	} else {
		b = b.Order(DefaultBannerOrder)
	}

	b = b.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := b.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type ExchangePager struct {
	Order  exchange.OrderOption
	Filter func(*ExchangeQuery) (*ExchangeQuery, error)
}

// ExchangePaginateOption enables pagination customization.
type ExchangePaginateOption func(*ExchangePager)

// DefaultExchangeOrder is the default ordering of Exchange.
var DefaultExchangeOrder = Desc(exchange.FieldID)

func newExchangePager(opts []ExchangePaginateOption) (*ExchangePager, error) {
	pager := &ExchangePager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultExchangeOrder
	}
	return pager, nil
}

func (p *ExchangePager) ApplyFilter(query *ExchangeQuery) (*ExchangeQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// ExchangePageList is Exchange PageList result.
type ExchangePageList struct {
	List        []*Exchange  `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (e *ExchangeQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...ExchangePaginateOption,
) (*ExchangePageList, error) {

	pager, err := newExchangePager(opts)
	if err != nil {
		return nil, err
	}

	if e, err = pager.ApplyFilter(e); err != nil {
		return nil, err
	}

	ret := &ExchangePageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := e.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		e = e.Order(pager.Order)
	} else {
		e = e.Order(DefaultExchangeOrder)
	}

	e = e.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := e.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type FilePager struct {
	Order  file.OrderOption
	Filter func(*FileQuery) (*FileQuery, error)
}

// FilePaginateOption enables pagination customization.
type FilePaginateOption func(*FilePager)

// DefaultFileOrder is the default ordering of File.
var DefaultFileOrder = Desc(file.FieldID)

func newFilePager(opts []FilePaginateOption) (*FilePager, error) {
	pager := &FilePager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultFileOrder
	}
	return pager, nil
}

func (p *FilePager) ApplyFilter(query *FileQuery) (*FileQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// FilePageList is File PageList result.
type FilePageList struct {
	List        []*File      `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (f *FileQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...FilePaginateOption,
) (*FilePageList, error) {

	pager, err := newFilePager(opts)
	if err != nil {
		return nil, err
	}

	if f, err = pager.ApplyFilter(f); err != nil {
		return nil, err
	}

	ret := &FilePageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := f.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		f = f.Order(pager.Order)
	} else {
		f = f.Order(DefaultFileOrder)
	}

	f = f.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := f.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type OrderPager struct {
	Order  order.OrderOption
	Filter func(*OrderQuery) (*OrderQuery, error)
}

// OrderPaginateOption enables pagination customization.
type OrderPaginateOption func(*OrderPager)

// DefaultOrderOrder is the default ordering of Order.
var DefaultOrderOrder = Desc(order.FieldID)

func newOrderPager(opts []OrderPaginateOption) (*OrderPager, error) {
	pager := &OrderPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultOrderOrder
	}
	return pager, nil
}

func (p *OrderPager) ApplyFilter(query *OrderQuery) (*OrderQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// OrderPageList is Order PageList result.
type OrderPageList struct {
	List        []*Order     `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (o *OrderQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...OrderPaginateOption,
) (*OrderPageList, error) {

	pager, err := newOrderPager(opts)
	if err != nil {
		return nil, err
	}

	if o, err = pager.ApplyFilter(o); err != nil {
		return nil, err
	}

	ret := &OrderPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := o.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		o = o.Order(pager.Order)
	} else {
		o = o.Order(DefaultOrderOrder)
	}

	o = o.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := o.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type OriginInviteCodePager struct {
	Order  origininvitecode.OrderOption
	Filter func(*OriginInviteCodeQuery) (*OriginInviteCodeQuery, error)
}

// OriginInviteCodePaginateOption enables pagination customization.
type OriginInviteCodePaginateOption func(*OriginInviteCodePager)

// DefaultOriginInviteCodeOrder is the default ordering of OriginInviteCode.
var DefaultOriginInviteCodeOrder = Desc(origininvitecode.FieldID)

func newOriginInviteCodePager(opts []OriginInviteCodePaginateOption) (*OriginInviteCodePager, error) {
	pager := &OriginInviteCodePager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultOriginInviteCodeOrder
	}
	return pager, nil
}

func (p *OriginInviteCodePager) ApplyFilter(query *OriginInviteCodeQuery) (*OriginInviteCodeQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// OriginInviteCodePageList is OriginInviteCode PageList result.
type OriginInviteCodePageList struct {
	List        []*OriginInviteCode `json:"list"`
	PageDetails *PageDetails        `json:"pageDetails"`
}

func (oic *OriginInviteCodeQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...OriginInviteCodePaginateOption,
) (*OriginInviteCodePageList, error) {

	pager, err := newOriginInviteCodePager(opts)
	if err != nil {
		return nil, err
	}

	if oic, err = pager.ApplyFilter(oic); err != nil {
		return nil, err
	}

	ret := &OriginInviteCodePageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := oic.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		oic = oic.Order(pager.Order)
	} else {
		oic = oic.Order(DefaultOriginInviteCodeOrder)
	}

	oic = oic.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := oic.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type PlayerPager struct {
	Order  player.OrderOption
	Filter func(*PlayerQuery) (*PlayerQuery, error)
}

// PlayerPaginateOption enables pagination customization.
type PlayerPaginateOption func(*PlayerPager)

// DefaultPlayerOrder is the default ordering of Player.
var DefaultPlayerOrder = Desc(player.FieldID)

func newPlayerPager(opts []PlayerPaginateOption) (*PlayerPager, error) {
	pager := &PlayerPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultPlayerOrder
	}
	return pager, nil
}

func (p *PlayerPager) ApplyFilter(query *PlayerQuery) (*PlayerQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// PlayerPageList is Player PageList result.
type PlayerPageList struct {
	List        []*Player    `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (pl *PlayerQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...PlayerPaginateOption,
) (*PlayerPageList, error) {

	pager, err := newPlayerPager(opts)
	if err != nil {
		return nil, err
	}

	if pl, err = pager.ApplyFilter(pl); err != nil {
		return nil, err
	}

	ret := &PlayerPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := pl.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		pl = pl.Order(pager.Order)
	} else {
		pl = pl.Order(DefaultPlayerOrder)
	}

	pl = pl.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := pl.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type PoolPager struct {
	Order  pool.OrderOption
	Filter func(*PoolQuery) (*PoolQuery, error)
}

// PoolPaginateOption enables pagination customization.
type PoolPaginateOption func(*PoolPager)

// DefaultPoolOrder is the default ordering of Pool.
var DefaultPoolOrder = Desc(pool.FieldID)

func newPoolPager(opts []PoolPaginateOption) (*PoolPager, error) {
	pager := &PoolPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultPoolOrder
	}
	return pager, nil
}

func (p *PoolPager) ApplyFilter(query *PoolQuery) (*PoolQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// PoolPageList is Pool PageList result.
type PoolPageList struct {
	List        []*Pool      `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (po *PoolQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...PoolPaginateOption,
) (*PoolPageList, error) {

	pager, err := newPoolPager(opts)
	if err != nil {
		return nil, err
	}

	if po, err = pager.ApplyFilter(po); err != nil {
		return nil, err
	}

	ret := &PoolPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := po.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		po = po.Order(pager.Order)
	} else {
		po = po.Order(DefaultPoolOrder)
	}

	po = po.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := po.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type RewardPager struct {
	Order  reward.OrderOption
	Filter func(*RewardQuery) (*RewardQuery, error)
}

// RewardPaginateOption enables pagination customization.
type RewardPaginateOption func(*RewardPager)

// DefaultRewardOrder is the default ordering of Reward.
var DefaultRewardOrder = Desc(reward.FieldID)

func newRewardPager(opts []RewardPaginateOption) (*RewardPager, error) {
	pager := &RewardPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultRewardOrder
	}
	return pager, nil
}

func (p *RewardPager) ApplyFilter(query *RewardQuery) (*RewardQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// RewardPageList is Reward PageList result.
type RewardPageList struct {
	List        []*Reward    `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (r *RewardQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...RewardPaginateOption,
) (*RewardPageList, error) {

	pager, err := newRewardPager(opts)
	if err != nil {
		return nil, err
	}

	if r, err = pager.ApplyFilter(r); err != nil {
		return nil, err
	}

	ret := &RewardPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := r.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		r = r.Order(pager.Order)
	} else {
		r = r.Order(DefaultRewardOrder)
	}

	r = r.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type RoundPager struct {
	Order  round.OrderOption
	Filter func(*RoundQuery) (*RoundQuery, error)
}

// RoundPaginateOption enables pagination customization.
type RoundPaginateOption func(*RoundPager)

// DefaultRoundOrder is the default ordering of Round.
var DefaultRoundOrder = Desc(round.FieldID)

func newRoundPager(opts []RoundPaginateOption) (*RoundPager, error) {
	pager := &RoundPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultRoundOrder
	}
	return pager, nil
}

func (p *RoundPager) ApplyFilter(query *RoundQuery) (*RoundQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// RoundPageList is Round PageList result.
type RoundPageList struct {
	List        []*Round     `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (r *RoundQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...RoundPaginateOption,
) (*RoundPageList, error) {

	pager, err := newRoundPager(opts)
	if err != nil {
		return nil, err
	}

	if r, err = pager.ApplyFilter(r); err != nil {
		return nil, err
	}

	ret := &RoundPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := r.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		r = r.Order(pager.Order)
	} else {
		r = r.Order(DefaultRoundOrder)
	}

	r = r.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type RoundInvestPager struct {
	Order  roundinvest.OrderOption
	Filter func(*RoundInvestQuery) (*RoundInvestQuery, error)
}

// RoundInvestPaginateOption enables pagination customization.
type RoundInvestPaginateOption func(*RoundInvestPager)

// DefaultRoundInvestOrder is the default ordering of RoundInvest.
var DefaultRoundInvestOrder = Desc(roundinvest.FieldID)

func newRoundInvestPager(opts []RoundInvestPaginateOption) (*RoundInvestPager, error) {
	pager := &RoundInvestPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultRoundInvestOrder
	}
	return pager, nil
}

func (p *RoundInvestPager) ApplyFilter(query *RoundInvestQuery) (*RoundInvestQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// RoundInvestPageList is RoundInvest PageList result.
type RoundInvestPageList struct {
	List        []*RoundInvest `json:"list"`
	PageDetails *PageDetails   `json:"pageDetails"`
}

func (ri *RoundInvestQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...RoundInvestPaginateOption,
) (*RoundInvestPageList, error) {

	pager, err := newRoundInvestPager(opts)
	if err != nil {
		return nil, err
	}

	if ri, err = pager.ApplyFilter(ri); err != nil {
		return nil, err
	}

	ret := &RoundInvestPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := ri.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		ri = ri.Order(pager.Order)
	} else {
		ri = ri.Order(DefaultRoundInvestOrder)
	}

	ri = ri.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := ri.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type RoundLambFoldPager struct {
	Order  roundlambfold.OrderOption
	Filter func(*RoundLambFoldQuery) (*RoundLambFoldQuery, error)
}

// RoundLambFoldPaginateOption enables pagination customization.
type RoundLambFoldPaginateOption func(*RoundLambFoldPager)

// DefaultRoundLambFoldOrder is the default ordering of RoundLambFold.
var DefaultRoundLambFoldOrder = Desc(roundlambfold.FieldID)

func newRoundLambFoldPager(opts []RoundLambFoldPaginateOption) (*RoundLambFoldPager, error) {
	pager := &RoundLambFoldPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultRoundLambFoldOrder
	}
	return pager, nil
}

func (p *RoundLambFoldPager) ApplyFilter(query *RoundLambFoldQuery) (*RoundLambFoldQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// RoundLambFoldPageList is RoundLambFold PageList result.
type RoundLambFoldPageList struct {
	List        []*RoundLambFold `json:"list"`
	PageDetails *PageDetails     `json:"pageDetails"`
}

func (rlf *RoundLambFoldQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...RoundLambFoldPaginateOption,
) (*RoundLambFoldPageList, error) {

	pager, err := newRoundLambFoldPager(opts)
	if err != nil {
		return nil, err
	}

	if rlf, err = pager.ApplyFilter(rlf); err != nil {
		return nil, err
	}

	ret := &RoundLambFoldPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := rlf.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		rlf = rlf.Order(pager.Order)
	} else {
		rlf = rlf.Order(DefaultRoundLambFoldOrder)
	}

	rlf = rlf.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := rlf.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type SettingPager struct {
	Order  setting.OrderOption
	Filter func(*SettingQuery) (*SettingQuery, error)
}

// SettingPaginateOption enables pagination customization.
type SettingPaginateOption func(*SettingPager)

// DefaultSettingOrder is the default ordering of Setting.
var DefaultSettingOrder = Desc(setting.FieldID)

func newSettingPager(opts []SettingPaginateOption) (*SettingPager, error) {
	pager := &SettingPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultSettingOrder
	}
	return pager, nil
}

func (p *SettingPager) ApplyFilter(query *SettingQuery) (*SettingQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// SettingPageList is Setting PageList result.
type SettingPageList struct {
	List        []*Setting   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (s *SettingQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...SettingPaginateOption,
) (*SettingPageList, error) {

	pager, err := newSettingPager(opts)
	if err != nil {
		return nil, err
	}

	if s, err = pager.ApplyFilter(s); err != nil {
		return nil, err
	}

	ret := &SettingPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := s.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		s = s.Order(pager.Order)
	} else {
		s = s.Order(DefaultSettingOrder)
	}

	s = s.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type StatementPager struct {
	Order  statement.OrderOption
	Filter func(*StatementQuery) (*StatementQuery, error)
}

// StatementPaginateOption enables pagination customization.
type StatementPaginateOption func(*StatementPager)

// DefaultStatementOrder is the default ordering of Statement.
var DefaultStatementOrder = Desc(statement.FieldID)

func newStatementPager(opts []StatementPaginateOption) (*StatementPager, error) {
	pager := &StatementPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultStatementOrder
	}
	return pager, nil
}

func (p *StatementPager) ApplyFilter(query *StatementQuery) (*StatementQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// StatementPageList is Statement PageList result.
type StatementPageList struct {
	List        []*Statement `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (s *StatementQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...StatementPaginateOption,
) (*StatementPageList, error) {

	pager, err := newStatementPager(opts)
	if err != nil {
		return nil, err
	}

	if s, err = pager.ApplyFilter(s); err != nil {
		return nil, err
	}

	ret := &StatementPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	query := s.Clone()
	query.ctx.Fields = nil
	count, err := query.Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		s = s.Order(pager.Order)
	} else {
		s = s.Order(DefaultStatementOrder)
	}

	s = s.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
