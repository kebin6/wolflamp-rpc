// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/kebin6/wolflamp-rpc/ent/banner"
	"github.com/kebin6/wolflamp-rpc/ent/exchange"
	"github.com/kebin6/wolflamp-rpc/ent/file"
	"github.com/kebin6/wolflamp-rpc/ent/order"
	"github.com/kebin6/wolflamp-rpc/ent/origininvitecode"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/ent/reward"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
	"github.com/kebin6/wolflamp-rpc/ent/schema"
	"github.com/kebin6/wolflamp-rpc/ent/setting"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bannerMixin := schema.Banner{}.Mixin()
	bannerMixinFields0 := bannerMixin[0].Fields()
	_ = bannerMixinFields0
	bannerMixinFields1 := bannerMixin[1].Fields()
	_ = bannerMixinFields1
	bannerFields := schema.Banner{}.Fields()
	_ = bannerFields
	// bannerDescCreatedAt is the schema descriptor for created_at field.
	bannerDescCreatedAt := bannerMixinFields0[1].Descriptor()
	// banner.DefaultCreatedAt holds the default value on creation for the created_at field.
	banner.DefaultCreatedAt = bannerDescCreatedAt.Default.(func() time.Time)
	// bannerDescUpdatedAt is the schema descriptor for updated_at field.
	bannerDescUpdatedAt := bannerMixinFields0[2].Descriptor()
	// banner.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	banner.DefaultUpdatedAt = bannerDescUpdatedAt.Default.(func() time.Time)
	// banner.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	banner.UpdateDefaultUpdatedAt = bannerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// bannerDescStatus is the schema descriptor for status field.
	bannerDescStatus := bannerMixinFields1[0].Descriptor()
	// banner.DefaultStatus holds the default value on creation for the status field.
	banner.DefaultStatus = bannerDescStatus.Default.(uint8)
	// bannerDescEndpoint is the schema descriptor for endpoint field.
	bannerDescEndpoint := bannerFields[0].Descriptor()
	// banner.DefaultEndpoint holds the default value on creation for the endpoint field.
	banner.DefaultEndpoint = bannerDescEndpoint.Default.(string)
	// bannerDescModule is the schema descriptor for module field.
	bannerDescModule := bannerFields[1].Descriptor()
	// banner.DefaultModule holds the default value on creation for the module field.
	banner.DefaultModule = bannerDescModule.Default.(string)
	// bannerDescFileType is the schema descriptor for file_type field.
	bannerDescFileType := bannerFields[2].Descriptor()
	// banner.DefaultFileType holds the default value on creation for the file_type field.
	banner.DefaultFileType = bannerDescFileType.Default.(uint8)
	// bannerDescPath is the schema descriptor for path field.
	bannerDescPath := bannerFields[3].Descriptor()
	// banner.DefaultPath holds the default value on creation for the path field.
	banner.DefaultPath = bannerDescPath.Default.(string)
	// bannerDescJumpURL is the schema descriptor for jump_url field.
	bannerDescJumpURL := bannerFields[4].Descriptor()
	// banner.DefaultJumpURL holds the default value on creation for the jump_url field.
	banner.DefaultJumpURL = bannerDescJumpURL.Default.(string)
	exchangeMixin := schema.Exchange{}.Mixin()
	exchangeMixinFields0 := exchangeMixin[0].Fields()
	_ = exchangeMixinFields0
	exchangeMixinFields1 := exchangeMixin[1].Fields()
	_ = exchangeMixinFields1
	exchangeFields := schema.Exchange{}.Fields()
	_ = exchangeFields
	// exchangeDescCreatedAt is the schema descriptor for created_at field.
	exchangeDescCreatedAt := exchangeMixinFields0[1].Descriptor()
	// exchange.DefaultCreatedAt holds the default value on creation for the created_at field.
	exchange.DefaultCreatedAt = exchangeDescCreatedAt.Default.(func() time.Time)
	// exchangeDescUpdatedAt is the schema descriptor for updated_at field.
	exchangeDescUpdatedAt := exchangeMixinFields0[2].Descriptor()
	// exchange.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	exchange.DefaultUpdatedAt = exchangeDescUpdatedAt.Default.(func() time.Time)
	// exchange.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	exchange.UpdateDefaultUpdatedAt = exchangeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// exchangeDescStatus is the schema descriptor for status field.
	exchangeDescStatus := exchangeMixinFields1[0].Descriptor()
	// exchange.DefaultStatus holds the default value on creation for the status field.
	exchange.DefaultStatus = exchangeDescStatus.Default.(uint8)
	// exchangeDescPlayerID is the schema descriptor for player_id field.
	exchangeDescPlayerID := exchangeFields[0].Descriptor()
	// exchange.DefaultPlayerID holds the default value on creation for the player_id field.
	exchange.DefaultPlayerID = exchangeDescPlayerID.Default.(uint64)
	// exchangeDescTransactionID is the schema descriptor for transaction_id field.
	exchangeDescTransactionID := exchangeFields[1].Descriptor()
	// exchange.DefaultTransactionID holds the default value on creation for the transaction_id field.
	exchange.DefaultTransactionID = exchangeDescTransactionID.Default.(string)
	// exchangeDescType is the schema descriptor for type field.
	exchangeDescType := exchangeFields[2].Descriptor()
	// exchange.DefaultType holds the default value on creation for the type field.
	exchange.DefaultType = exchangeDescType.Default.(uint32)
	// exchangeDescCoinNum is the schema descriptor for coin_num field.
	exchangeDescCoinNum := exchangeFields[3].Descriptor()
	// exchange.DefaultCoinNum holds the default value on creation for the coin_num field.
	exchange.DefaultCoinNum = exchangeDescCoinNum.Default.(uint32)
	// exchangeDescLampNum is the schema descriptor for lamp_num field.
	exchangeDescLampNum := exchangeFields[4].Descriptor()
	// exchange.DefaultLampNum holds the default value on creation for the lamp_num field.
	exchange.DefaultLampNum = exchangeDescLampNum.Default.(uint32)
	fileMixin := schema.File{}.Mixin()
	fileMixinFields0 := fileMixin[0].Fields()
	_ = fileMixinFields0
	fileMixinFields1 := fileMixin[1].Fields()
	_ = fileMixinFields1
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescCreatedAt is the schema descriptor for created_at field.
	fileDescCreatedAt := fileMixinFields0[1].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the created_at field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	// fileDescUpdatedAt is the schema descriptor for updated_at field.
	fileDescUpdatedAt := fileMixinFields0[2].Descriptor()
	// file.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	file.DefaultUpdatedAt = fileDescUpdatedAt.Default.(func() time.Time)
	// file.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	file.UpdateDefaultUpdatedAt = fileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// fileDescStatus is the schema descriptor for status field.
	fileDescStatus := fileMixinFields1[0].Descriptor()
	// file.DefaultStatus holds the default value on creation for the status field.
	file.DefaultStatus = fileDescStatus.Default.(uint8)
	// fileDescID is the schema descriptor for id field.
	fileDescID := fileMixinFields0[0].Descriptor()
	// file.DefaultID holds the default value on creation for the id field.
	file.DefaultID = fileDescID.Default.(func() uuid.UUID)
	orderMixin := schema.Order{}.Mixin()
	orderMixinFields0 := orderMixin[0].Fields()
	_ = orderMixinFields0
	orderMixinFields1 := orderMixin[1].Fields()
	_ = orderMixinFields1
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderMixinFields0[1].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() time.Time)
	// orderDescUpdatedAt is the schema descriptor for updated_at field.
	orderDescUpdatedAt := orderMixinFields0[2].Descriptor()
	// order.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	order.DefaultUpdatedAt = orderDescUpdatedAt.Default.(func() time.Time)
	// order.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	order.UpdateDefaultUpdatedAt = orderDescUpdatedAt.UpdateDefault.(func() time.Time)
	// orderDescStatus is the schema descriptor for status field.
	orderDescStatus := orderMixinFields1[0].Descriptor()
	// order.DefaultStatus holds the default value on creation for the status field.
	order.DefaultStatus = orderDescStatus.Default.(uint8)
	// orderDescPlayerID is the schema descriptor for player_id field.
	orderDescPlayerID := orderFields[0].Descriptor()
	// order.DefaultPlayerID holds the default value on creation for the player_id field.
	order.DefaultPlayerID = orderDescPlayerID.Default.(uint64)
	// orderDescType is the schema descriptor for type field.
	orderDescType := orderFields[1].Descriptor()
	// order.DefaultType holds the default value on creation for the type field.
	order.DefaultType = orderDescType.Default.(string)
	// orderDescCode is the schema descriptor for code field.
	orderDescCode := orderFields[2].Descriptor()
	// order.DefaultCode holds the default value on creation for the code field.
	order.DefaultCode = orderDescCode.Default.(string)
	// orderDescTransactionID is the schema descriptor for transaction_id field.
	orderDescTransactionID := orderFields[3].Descriptor()
	// order.DefaultTransactionID holds the default value on creation for the transaction_id field.
	order.DefaultTransactionID = orderDescTransactionID.Default.(string)
	// orderDescFromAddress is the schema descriptor for from_address field.
	orderDescFromAddress := orderFields[4].Descriptor()
	// order.DefaultFromAddress holds the default value on creation for the from_address field.
	order.DefaultFromAddress = orderDescFromAddress.Default.(string)
	// orderDescToAddress is the schema descriptor for to_address field.
	orderDescToAddress := orderFields[5].Descriptor()
	// order.DefaultToAddress holds the default value on creation for the to_address field.
	order.DefaultToAddress = orderDescToAddress.Default.(string)
	// orderDescNum is the schema descriptor for num field.
	orderDescNum := orderFields[6].Descriptor()
	// order.DefaultNum holds the default value on creation for the num field.
	order.DefaultNum = orderDescNum.Default.(float64)
	// orderDescHandlingFee is the schema descriptor for handling_fee field.
	orderDescHandlingFee := orderFields[7].Descriptor()
	// order.DefaultHandlingFee holds the default value on creation for the handling_fee field.
	order.DefaultHandlingFee = orderDescHandlingFee.Default.(float64)
	// orderDescNetwork is the schema descriptor for network field.
	orderDescNetwork := orderFields[8].Descriptor()
	// order.DefaultNetwork holds the default value on creation for the network field.
	order.DefaultNetwork = orderDescNetwork.Default.(string)
	// orderDescRemark is the schema descriptor for remark field.
	orderDescRemark := orderFields[9].Descriptor()
	// order.DefaultRemark holds the default value on creation for the remark field.
	order.DefaultRemark = orderDescRemark.Default.(string)
	origininvitecodeMixin := schema.OriginInviteCode{}.Mixin()
	origininvitecodeMixinFields0 := origininvitecodeMixin[0].Fields()
	_ = origininvitecodeMixinFields0
	origininvitecodeMixinFields1 := origininvitecodeMixin[1].Fields()
	_ = origininvitecodeMixinFields1
	origininvitecodeFields := schema.OriginInviteCode{}.Fields()
	_ = origininvitecodeFields
	// origininvitecodeDescCreatedAt is the schema descriptor for created_at field.
	origininvitecodeDescCreatedAt := origininvitecodeMixinFields0[1].Descriptor()
	// origininvitecode.DefaultCreatedAt holds the default value on creation for the created_at field.
	origininvitecode.DefaultCreatedAt = origininvitecodeDescCreatedAt.Default.(func() time.Time)
	// origininvitecodeDescUpdatedAt is the schema descriptor for updated_at field.
	origininvitecodeDescUpdatedAt := origininvitecodeMixinFields0[2].Descriptor()
	// origininvitecode.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	origininvitecode.DefaultUpdatedAt = origininvitecodeDescUpdatedAt.Default.(func() time.Time)
	// origininvitecode.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	origininvitecode.UpdateDefaultUpdatedAt = origininvitecodeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// origininvitecodeDescStatus is the schema descriptor for status field.
	origininvitecodeDescStatus := origininvitecodeMixinFields1[0].Descriptor()
	// origininvitecode.DefaultStatus holds the default value on creation for the status field.
	origininvitecode.DefaultStatus = origininvitecodeDescStatus.Default.(uint8)
	// origininvitecodeDescCode is the schema descriptor for code field.
	origininvitecodeDescCode := origininvitecodeFields[0].Descriptor()
	// origininvitecode.DefaultCode holds the default value on creation for the code field.
	origininvitecode.DefaultCode = origininvitecodeDescCode.Default.(string)
	playerMixin := schema.Player{}.Mixin()
	playerMixinFields0 := playerMixin[0].Fields()
	_ = playerMixinFields0
	playerMixinFields1 := playerMixin[1].Fields()
	_ = playerMixinFields1
	playerFields := schema.Player{}.Fields()
	_ = playerFields
	// playerDescCreatedAt is the schema descriptor for created_at field.
	playerDescCreatedAt := playerMixinFields0[1].Descriptor()
	// player.DefaultCreatedAt holds the default value on creation for the created_at field.
	player.DefaultCreatedAt = playerDescCreatedAt.Default.(func() time.Time)
	// playerDescUpdatedAt is the schema descriptor for updated_at field.
	playerDescUpdatedAt := playerMixinFields0[2].Descriptor()
	// player.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	player.DefaultUpdatedAt = playerDescUpdatedAt.Default.(func() time.Time)
	// player.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	player.UpdateDefaultUpdatedAt = playerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// playerDescStatus is the schema descriptor for status field.
	playerDescStatus := playerMixinFields1[0].Descriptor()
	// player.DefaultStatus holds the default value on creation for the status field.
	player.DefaultStatus = playerDescStatus.Default.(uint8)
	// playerDescName is the schema descriptor for name field.
	playerDescName := playerFields[0].Descriptor()
	// player.DefaultName holds the default value on creation for the name field.
	player.DefaultName = playerDescName.Default.(string)
	// playerDescEmail is the schema descriptor for email field.
	playerDescEmail := playerFields[1].Descriptor()
	// player.DefaultEmail holds the default value on creation for the email field.
	player.DefaultEmail = playerDescEmail.Default.(string)
	// playerDescPassword is the schema descriptor for password field.
	playerDescPassword := playerFields[2].Descriptor()
	// player.DefaultPassword holds the default value on creation for the password field.
	player.DefaultPassword = playerDescPassword.Default.(string)
	// playerDescTransactionPassword is the schema descriptor for transaction_password field.
	playerDescTransactionPassword := playerFields[3].Descriptor()
	// player.DefaultTransactionPassword holds the default value on creation for the transaction_password field.
	player.DefaultTransactionPassword = playerDescTransactionPassword.Default.(string)
	// playerDescLamp is the schema descriptor for lamp field.
	playerDescLamp := playerFields[4].Descriptor()
	// player.DefaultLamp holds the default value on creation for the lamp field.
	player.DefaultLamp = playerDescLamp.Default.(float32)
	// playerDescRank is the schema descriptor for rank field.
	playerDescRank := playerFields[5].Descriptor()
	// player.DefaultRank holds the default value on creation for the rank field.
	player.DefaultRank = playerDescRank.Default.(uint32)
	// playerDescAmount is the schema descriptor for amount field.
	playerDescAmount := playerFields[6].Descriptor()
	// player.DefaultAmount holds the default value on creation for the amount field.
	player.DefaultAmount = playerDescAmount.Default.(float64)
	// playerDescDepositAddress is the schema descriptor for deposit_address field.
	playerDescDepositAddress := playerFields[7].Descriptor()
	// player.DefaultDepositAddress holds the default value on creation for the deposit_address field.
	player.DefaultDepositAddress = playerDescDepositAddress.Default.(string)
	// playerDescInvitedNum is the schema descriptor for invited_num field.
	playerDescInvitedNum := playerFields[8].Descriptor()
	// player.DefaultInvitedNum holds the default value on creation for the invited_num field.
	player.DefaultInvitedNum = playerDescInvitedNum.Default.(uint32)
	// playerDescTotalIncome is the schema descriptor for total_income field.
	playerDescTotalIncome := playerFields[9].Descriptor()
	// player.DefaultTotalIncome holds the default value on creation for the total_income field.
	player.DefaultTotalIncome = playerDescTotalIncome.Default.(float64)
	// playerDescProfitAndLoss is the schema descriptor for profit_and_loss field.
	playerDescProfitAndLoss := playerFields[10].Descriptor()
	// player.DefaultProfitAndLoss holds the default value on creation for the profit_and_loss field.
	player.DefaultProfitAndLoss = playerDescProfitAndLoss.Default.(float32)
	// playerDescRecent100WinPercent is the schema descriptor for recent_100_win_percent field.
	playerDescRecent100WinPercent := playerFields[11].Descriptor()
	// player.DefaultRecent100WinPercent holds the default value on creation for the recent_100_win_percent field.
	player.DefaultRecent100WinPercent = playerDescRecent100WinPercent.Default.(float32)
	// playerDescInviteCode is the schema descriptor for invite_code field.
	playerDescInviteCode := playerFields[12].Descriptor()
	// player.DefaultInviteCode holds the default value on creation for the invite_code field.
	player.DefaultInviteCode = playerDescInviteCode.Default.(string)
	// playerDescInviterID is the schema descriptor for inviter_id field.
	playerDescInviterID := playerFields[13].Descriptor()
	// player.DefaultInviterID holds the default value on creation for the inviter_id field.
	player.DefaultInviterID = playerDescInviterID.Default.(uint64)
	// playerDescInvitedCode is the schema descriptor for invited_code field.
	playerDescInvitedCode := playerFields[14].Descriptor()
	// player.DefaultInvitedCode holds the default value on creation for the invited_code field.
	player.DefaultInvitedCode = playerDescInvitedCode.Default.(string)
	// playerDescSystemCommission is the schema descriptor for system_commission field.
	playerDescSystemCommission := playerFields[15].Descriptor()
	// player.DefaultSystemCommission holds the default value on creation for the system_commission field.
	player.DefaultSystemCommission = playerDescSystemCommission.Default.(float32)
	rewardMixin := schema.Reward{}.Mixin()
	rewardMixinFields0 := rewardMixin[0].Fields()
	_ = rewardMixinFields0
	rewardMixinFields1 := rewardMixin[1].Fields()
	_ = rewardMixinFields1
	rewardFields := schema.Reward{}.Fields()
	_ = rewardFields
	// rewardDescCreatedAt is the schema descriptor for created_at field.
	rewardDescCreatedAt := rewardMixinFields0[1].Descriptor()
	// reward.DefaultCreatedAt holds the default value on creation for the created_at field.
	reward.DefaultCreatedAt = rewardDescCreatedAt.Default.(func() time.Time)
	// rewardDescUpdatedAt is the schema descriptor for updated_at field.
	rewardDescUpdatedAt := rewardMixinFields0[2].Descriptor()
	// reward.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	reward.DefaultUpdatedAt = rewardDescUpdatedAt.Default.(func() time.Time)
	// reward.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	reward.UpdateDefaultUpdatedAt = rewardDescUpdatedAt.UpdateDefault.(func() time.Time)
	// rewardDescStatus is the schema descriptor for status field.
	rewardDescStatus := rewardMixinFields1[0].Descriptor()
	// reward.DefaultStatus holds the default value on creation for the status field.
	reward.DefaultStatus = rewardDescStatus.Default.(uint8)
	// rewardDescToID is the schema descriptor for to_id field.
	rewardDescToID := rewardFields[0].Descriptor()
	// reward.DefaultToID holds the default value on creation for the to_id field.
	reward.DefaultToID = rewardDescToID.Default.(uint64)
	// rewardDescContributorID is the schema descriptor for contributor_id field.
	rewardDescContributorID := rewardFields[1].Descriptor()
	// reward.DefaultContributorID holds the default value on creation for the contributor_id field.
	reward.DefaultContributorID = rewardDescContributorID.Default.(uint64)
	// rewardDescContributorEmail is the schema descriptor for contributor_email field.
	rewardDescContributorEmail := rewardFields[2].Descriptor()
	// reward.DefaultContributorEmail holds the default value on creation for the contributor_email field.
	reward.DefaultContributorEmail = rewardDescContributorEmail.Default.(string)
	// rewardDescContributorLevel is the schema descriptor for contributor_level field.
	rewardDescContributorLevel := rewardFields[3].Descriptor()
	// reward.DefaultContributorLevel holds the default value on creation for the contributor_level field.
	reward.DefaultContributorLevel = rewardDescContributorLevel.Default.(uint32)
	// rewardDescNum is the schema descriptor for num field.
	rewardDescNum := rewardFields[4].Descriptor()
	// reward.DefaultNum holds the default value on creation for the num field.
	reward.DefaultNum = rewardDescNum.Default.(float32)
	// rewardDescFormula is the schema descriptor for formula field.
	rewardDescFormula := rewardFields[5].Descriptor()
	// reward.DefaultFormula holds the default value on creation for the formula field.
	reward.DefaultFormula = rewardDescFormula.Default.(string)
	// rewardDescRemark is the schema descriptor for remark field.
	rewardDescRemark := rewardFields[6].Descriptor()
	// reward.DefaultRemark holds the default value on creation for the remark field.
	reward.DefaultRemark = rewardDescRemark.Default.(string)
	roundMixin := schema.Round{}.Mixin()
	roundMixinFields0 := roundMixin[0].Fields()
	_ = roundMixinFields0
	roundMixinFields1 := roundMixin[1].Fields()
	_ = roundMixinFields1
	roundFields := schema.Round{}.Fields()
	_ = roundFields
	// roundDescCreatedAt is the schema descriptor for created_at field.
	roundDescCreatedAt := roundMixinFields0[1].Descriptor()
	// round.DefaultCreatedAt holds the default value on creation for the created_at field.
	round.DefaultCreatedAt = roundDescCreatedAt.Default.(func() time.Time)
	// roundDescUpdatedAt is the schema descriptor for updated_at field.
	roundDescUpdatedAt := roundMixinFields0[2].Descriptor()
	// round.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	round.DefaultUpdatedAt = roundDescUpdatedAt.Default.(func() time.Time)
	// round.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	round.UpdateDefaultUpdatedAt = roundDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roundDescStatus is the schema descriptor for status field.
	roundDescStatus := roundMixinFields1[0].Descriptor()
	// round.DefaultStatus holds the default value on creation for the status field.
	round.DefaultStatus = roundDescStatus.Default.(uint8)
	// roundDescSelectedFold is the schema descriptor for selected_fold field.
	roundDescSelectedFold := roundFields[4].Descriptor()
	// round.DefaultSelectedFold holds the default value on creation for the selected_fold field.
	round.DefaultSelectedFold = roundDescSelectedFold.Default.(uint32)
	roundinvestMixin := schema.RoundInvest{}.Mixin()
	roundinvestMixinFields0 := roundinvestMixin[0].Fields()
	_ = roundinvestMixinFields0
	roundinvestFields := schema.RoundInvest{}.Fields()
	_ = roundinvestFields
	// roundinvestDescCreatedAt is the schema descriptor for created_at field.
	roundinvestDescCreatedAt := roundinvestMixinFields0[1].Descriptor()
	// roundinvest.DefaultCreatedAt holds the default value on creation for the created_at field.
	roundinvest.DefaultCreatedAt = roundinvestDescCreatedAt.Default.(func() time.Time)
	// roundinvestDescUpdatedAt is the schema descriptor for updated_at field.
	roundinvestDescUpdatedAt := roundinvestMixinFields0[2].Descriptor()
	// roundinvest.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	roundinvest.DefaultUpdatedAt = roundinvestDescUpdatedAt.Default.(func() time.Time)
	// roundinvest.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	roundinvest.UpdateDefaultUpdatedAt = roundinvestDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roundinvestDescPlayerID is the schema descriptor for player_id field.
	roundinvestDescPlayerID := roundinvestFields[0].Descriptor()
	// roundinvest.DefaultPlayerID holds the default value on creation for the player_id field.
	roundinvest.DefaultPlayerID = roundinvestDescPlayerID.Default.(uint64)
	// roundinvestDescPlayerEmail is the schema descriptor for player_email field.
	roundinvestDescPlayerEmail := roundinvestFields[1].Descriptor()
	// roundinvest.DefaultPlayerEmail holds the default value on creation for the player_email field.
	roundinvest.DefaultPlayerEmail = roundinvestDescPlayerEmail.Default.(string)
	// roundinvestDescFoldNo is the schema descriptor for fold_no field.
	roundinvestDescFoldNo := roundinvestFields[2].Descriptor()
	// roundinvest.DefaultFoldNo holds the default value on creation for the fold_no field.
	roundinvest.DefaultFoldNo = roundinvestDescFoldNo.Default.(uint32)
	// roundinvestDescLambNum is the schema descriptor for lamb_num field.
	roundinvestDescLambNum := roundinvestFields[3].Descriptor()
	// roundinvest.DefaultLambNum holds the default value on creation for the lamb_num field.
	roundinvest.DefaultLambNum = roundinvestDescLambNum.Default.(uint32)
	// roundinvestDescProfitAndLoss is the schema descriptor for profit_and_loss field.
	roundinvestDescProfitAndLoss := roundinvestFields[4].Descriptor()
	// roundinvest.DefaultProfitAndLoss holds the default value on creation for the profit_and_loss field.
	roundinvest.DefaultProfitAndLoss = roundinvestDescProfitAndLoss.Default.(float32)
	// roundinvestDescRoundID is the schema descriptor for round_id field.
	roundinvestDescRoundID := roundinvestFields[5].Descriptor()
	// roundinvest.DefaultRoundID holds the default value on creation for the round_id field.
	roundinvest.DefaultRoundID = roundinvestDescRoundID.Default.(uint64)
	// roundinvestDescRoundCount is the schema descriptor for round_count field.
	roundinvestDescRoundCount := roundinvestFields[6].Descriptor()
	// roundinvest.DefaultRoundCount holds the default value on creation for the round_count field.
	roundinvest.DefaultRoundCount = roundinvestDescRoundCount.Default.(uint32)
	roundlambfoldMixin := schema.RoundLambFold{}.Mixin()
	roundlambfoldMixinFields0 := roundlambfoldMixin[0].Fields()
	_ = roundlambfoldMixinFields0
	roundlambfoldFields := schema.RoundLambFold{}.Fields()
	_ = roundlambfoldFields
	// roundlambfoldDescCreatedAt is the schema descriptor for created_at field.
	roundlambfoldDescCreatedAt := roundlambfoldMixinFields0[1].Descriptor()
	// roundlambfold.DefaultCreatedAt holds the default value on creation for the created_at field.
	roundlambfold.DefaultCreatedAt = roundlambfoldDescCreatedAt.Default.(func() time.Time)
	// roundlambfoldDescUpdatedAt is the schema descriptor for updated_at field.
	roundlambfoldDescUpdatedAt := roundlambfoldMixinFields0[2].Descriptor()
	// roundlambfold.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	roundlambfold.DefaultUpdatedAt = roundlambfoldDescUpdatedAt.Default.(func() time.Time)
	// roundlambfold.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	roundlambfold.UpdateDefaultUpdatedAt = roundlambfoldDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roundlambfoldDescFoldNo is the schema descriptor for fold_no field.
	roundlambfoldDescFoldNo := roundlambfoldFields[0].Descriptor()
	// roundlambfold.DefaultFoldNo holds the default value on creation for the fold_no field.
	roundlambfold.DefaultFoldNo = roundlambfoldDescFoldNo.Default.(uint32)
	// roundlambfoldDescLambNum is the schema descriptor for lamb_num field.
	roundlambfoldDescLambNum := roundlambfoldFields[1].Descriptor()
	// roundlambfold.DefaultLambNum holds the default value on creation for the lamb_num field.
	roundlambfold.DefaultLambNum = roundlambfoldDescLambNum.Default.(uint32)
	// roundlambfoldDescRoundID is the schema descriptor for round_id field.
	roundlambfoldDescRoundID := roundlambfoldFields[2].Descriptor()
	// roundlambfold.DefaultRoundID holds the default value on creation for the round_id field.
	roundlambfold.DefaultRoundID = roundlambfoldDescRoundID.Default.(uint64)
	// roundlambfoldDescRoundCount is the schema descriptor for round_count field.
	roundlambfoldDescRoundCount := roundlambfoldFields[3].Descriptor()
	// roundlambfold.DefaultRoundCount holds the default value on creation for the round_count field.
	roundlambfold.DefaultRoundCount = roundlambfoldDescRoundCount.Default.(uint32)
	// roundlambfoldDescProfitAndLoss is the schema descriptor for profit_and_loss field.
	roundlambfoldDescProfitAndLoss := roundlambfoldFields[4].Descriptor()
	// roundlambfold.DefaultProfitAndLoss holds the default value on creation for the profit_and_loss field.
	roundlambfold.DefaultProfitAndLoss = roundlambfoldDescProfitAndLoss.Default.(float32)
	settingMixin := schema.Setting{}.Mixin()
	settingMixinFields0 := settingMixin[0].Fields()
	_ = settingMixinFields0
	settingMixinFields1 := settingMixin[1].Fields()
	_ = settingMixinFields1
	settingFields := schema.Setting{}.Fields()
	_ = settingFields
	// settingDescCreatedAt is the schema descriptor for created_at field.
	settingDescCreatedAt := settingMixinFields0[1].Descriptor()
	// setting.DefaultCreatedAt holds the default value on creation for the created_at field.
	setting.DefaultCreatedAt = settingDescCreatedAt.Default.(func() time.Time)
	// settingDescUpdatedAt is the schema descriptor for updated_at field.
	settingDescUpdatedAt := settingMixinFields0[2].Descriptor()
	// setting.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	setting.DefaultUpdatedAt = settingDescUpdatedAt.Default.(func() time.Time)
	// setting.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	setting.UpdateDefaultUpdatedAt = settingDescUpdatedAt.UpdateDefault.(func() time.Time)
	// settingDescStatus is the schema descriptor for status field.
	settingDescStatus := settingMixinFields1[0].Descriptor()
	// setting.DefaultStatus holds the default value on creation for the status field.
	setting.DefaultStatus = settingDescStatus.Default.(uint8)
	// settingDescName is the schema descriptor for name field.
	settingDescName := settingFields[0].Descriptor()
	// setting.DefaultName holds the default value on creation for the name field.
	setting.DefaultName = settingDescName.Default.(string)
	// settingDescModule is the schema descriptor for module field.
	settingDescModule := settingFields[1].Descriptor()
	// setting.DefaultModule holds the default value on creation for the module field.
	setting.DefaultModule = settingDescModule.Default.(string)
	// settingDescValue is the schema descriptor for value field.
	settingDescValue := settingFields[2].Descriptor()
	// setting.DefaultValue holds the default value on creation for the value field.
	setting.DefaultValue = settingDescValue.Default.(string)
}
