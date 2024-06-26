package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	virtualgrouptypes "github.com/evmos/evmos/v12/x/virtualgroup/types"
	"github.com/zkMeLabs/mechain-storage-provider/cmd/utils"

	"github.com/urfave/cli/v2"
)

const swapInCommands = "SwapIn Commands"

const (
	ProcessStatus_Recovering = "Recovering"
	ProcessStatus_Verifying  = "Verifying"
	ProcessStatus_Complete   = "Successful"
)

var gvgIDFlag = &cli.Uint64Flag{
	Name:     "gvgId",
	Usage:    "assign global virtual group id",
	Aliases:  []string{"gid"},
	Required: true,
}

var vgfIDFlag = &cli.Uint64Flag{
	Name:     "vgf",
	Usage:    "assign global virtual group family id",
	Aliases:  []string{"f"},
	Required: true,
}

var targetSPIDFlag = &cli.Uint64Flag{
	Name:     "targetSP",
	Usage:    "assign target sp",
	Aliases:  []string{"sp"},
	Required: true,
}

var SwapInCmd = &cli.Command{
	Action: SwapInAction,
	Name:   "swapIn",
	Usage:  "Successor swap in GVG/VGF",
	Flags: []cli.Flag{
		utils.ConfigFileFlag,
		gvgIDFlag,
		vgfIDFlag,
		targetSPIDFlag,
	},
	Category: swapInCommands,
	Description: `You can use this command if you know that an sp is exiting and are ready to take over` +
		`This command is for the transaction that sends a Swap In the chain`,
}

var RecoverGVGCmd = &cli.Command{
	Action: RecoverGVGAction,
	Name:   "recover-gvg",
	Usage:  "recover object in gvg",
	Flags: []cli.Flag{
		utils.ConfigFileFlag,
		gvgIDFlag,
	},
	Category: swapInCommands,
	Description: `After determining a successful successor you can use the "recover-gvg" CMD` +
		`This command notifies sp and causes sp to recover resource` +
		`Resources include all resources that exit sp serves as the secondary sp in gvg`,
}

var RecoverVGFCmd = &cli.Command{
	Action: RecoverVGFAction,
	Name:   "recover-vgf",
	Usage:  "recover objects in vgf",
	Flags: []cli.Flag{
		utils.ConfigFileFlag,
		vgfIDFlag,
	},
	Category: swapInCommands,
	Description: `After determining a successful successor you can use the "recover-gvg" CMD` +
		`This command notifies sp and causes sp to recover resource` +
		`Resources include all resources that exit sp serves as the primary sp in gvg family`,
}

var CompleteSwapInCmd = &cli.Command{
	Action: CompleteSwapInAction,
	Name:   "completeSwapIn",
	Usage:  "complete swap in",
	Flags: []cli.Flag{
		utils.ConfigFileFlag,
		gvgIDFlag,
		vgfIDFlag,
	},
	Category:    swapInCommands,
	Description: `After confirming that the recover resource is complete, send a complete swap in to the chain using the completeSwapIn command`,
}

var CancelSwapInCmd = &cli.Command{
	Action: CancelSwapInAction,
	Name:   "cancelSwapIn",
	Usage:  "cancel swap in",
	Flags: []cli.Flag{
		utils.ConfigFileFlag,
		gvgIDFlag,
		vgfIDFlag,
	},
	Category: swapInCommands,
}

var QueryRecoverProcessCmd = &cli.Command{
	Action: QueryRecoverProcessAction,
	Name:   "query-recover-p",
	Usage:  "query recover process",
	Flags: []cli.Flag{
		utils.ConfigFileFlag,
		gvgIDFlag,
		vgfIDFlag,
	},
	Category:    swapInCommands,
	Description: `It is used to query the recovery resource progress. The progress is displayed in std and recover_process.json`,
}

var ListGlobalVirtualGroupsBySecondarySPCmd = &cli.Command{
	Action: ListGlobalVirtualGroupsBySecondarySPAction,
	Name:   "query-gvg-by-sp",
	Usage:  "get GlobalVirtualGroups List By SecondarySP",
	Flags: []cli.Flag{
		utils.ConfigFileFlag,
		targetSPIDFlag,
	},
	Category:    swapInCommands,
	Description: `get GlobalVirtualGroups List By SecondarySP`,
}

var ListVirtualGroupFamiliesBySpIDCmd = &cli.Command{
	Action: ListVirtualGroupFamiliesBySpIDAction,
	Name:   "query-vgf-by-sp",
	Usage:  "get VirtualGroupFamily List By SpID",
	Flags: []cli.Flag{
		utils.ConfigFileFlag,
		targetSPIDFlag,
	},
	Category:    swapInCommands,
	Description: `get VirtualGroupFamily List By SpID`,
}

type FailedRecoverObject struct {
	ObjectId        uint64 `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	VirtualGroupId  uint32 `protobuf:"varint,2,opt,name=virtual_group_id,json=virtualGroupId,proto3" json:"virtual_group_id,omitempty"`
	RedundancyIndex int32  `protobuf:"varint,3,opt,name=redundancy_index,json=redundancyIndex,proto3" json:"redundancy_index,omitempty"`
	RetryTime       int32  `protobuf:"varint,4,opt,name=retry_time,json=retryTime,proto3" json:"retry_time,omitempty"`
}

type ProcessResult struct {
	VirtualGroupId         uint32                 `json:"virtual_group_id"`
	VirtualGroupFamilyId   uint32                 `json:"virtual_group_family_id"`
	RedundancyIndex        int32                  `json:"redundancy_index"`
	StartAfter             uint64                 `json:"start_after"`
	Limit                  uint64                 `json:"limit"`
	Status                 string                 `json:"status"`
	ObjectCount            uint64                 `json:"object_count"`
	FailedObjectTotalCount uint64                 `json:"failed_object_total_count"`
	RecoverFailedObject    []*FailedRecoverObject `json:"recover_failed_object"`
}

func SwapInAction(ctx *cli.Context) error {
	cfg, err := utils.MakeConfig(ctx)
	if err != nil {
		println(err.Error())
		return err
	}

	targetSpID := ctx.Uint64(targetSPIDFlag.Name)
	gvgID := ctx.Uint64(gvgIDFlag.Name)
	gvgfID := ctx.Uint64(vgfIDFlag.Name)

	reserveSwapIn := &virtualgrouptypes.MsgReserveSwapIn{
		TargetSpId:                 uint32(targetSpID),
		GlobalVirtualGroupFamilyId: uint32(gvgfID),
		GlobalVirtualGroupId:       uint32(gvgID),
		StorageProvider:            cfg.SpAccount.SpOperatorAddress,
	}

	spClient := utils.MakeGfSpClient(cfg)
	tx, err := spClient.ReserveSwapIn(ctx.Context, reserveSwapIn)
	if err != nil {
		println(err.Error())
		return err
	}
	fmt.Printf("tx successfully! tx_hash:%s", tx)
	return nil
}

func CompleteSwapInAction(ctx *cli.Context) error {
	cfg, err := utils.MakeConfig(ctx)
	if err != nil {
		println(err.Error())
		return err
	}

	gvgID := ctx.Uint64(gvgIDFlag.Name)
	gvgfID := ctx.Uint64(vgfIDFlag.Name)

	completeSwapIn := &virtualgrouptypes.MsgCompleteSwapIn{
		GlobalVirtualGroupFamilyId: uint32(gvgfID),
		GlobalVirtualGroupId:       uint32(gvgID),
		StorageProvider:            cfg.SpAccount.SpOperatorAddress,
	}

	spClient := utils.MakeGfSpClient(cfg)
	tx, err := spClient.CompleteSwapIn(ctx.Context, completeSwapIn)
	if err != nil {
		println(err.Error())
		return err
	}
	fmt.Printf("tx successfully! tx_hash:%s", tx)
	return nil
}

func RecoverGVGAction(ctx *cli.Context) error {
	cfg, err := utils.MakeConfig(ctx)
	if err != nil {
		println(err.Error())
		return err
	}

	// get client
	chainClient, err := utils.MakeGnfd(cfg)
	if err != nil {
		println(err.Error())
		return err
	}
	spClient := utils.MakeGfSpClient(cfg)

	// check swapIn info
	sp, err := chainClient.QuerySP(ctx.Context, cfg.SpAccount.SpOperatorAddress)
	if err != nil {
		println(err.Error())
		return err
	}
	gvgID := ctx.Uint64(gvgIDFlag.Name)
	swapInInfo, err := chainClient.QuerySwapInInfo(ctx.Context, 0, uint32(gvgID))
	if err != nil {
		println(err.Error())
		return err
	}
	if swapInInfo.GetSuccessorSpId() != sp.GetId() {
		println("sp is not successor sp")
		return errors.New("sp is not successor sp")
	}

	// get replicateIndex
	gvgInfo, err := spClient.GetGlobalVirtualGroupByGvgID(ctx.Context, uint32(gvgID))
	if err != nil {
		println(err.Error())
		return err
	}
	var replicateIndex int32
	for idx, sspID := range gvgInfo.SecondarySpIds {
		if sspID == swapInInfo.TargetSpId {
			replicateIndex = int32(idx)
		}
	}

	_, executing, err := spClient.QueryRecoverProcess(ctx.Context, uint32(0), uint32(gvgID))
	if err != nil {
		println(err.Error())
		return err
	}
	if executing {
		println("Please wait until the previous recover work is completed")
		return nil
	}

	err = spClient.TriggerRecoverForSuccessorSP(ctx.Context, 0, uint32(gvgID), replicateIndex)
	if err != nil {
		println(err.Error())
		return err
	}
	println("Trigger successfully! please waiting process")
	return nil
}

func RecoverVGFAction(ctx *cli.Context) error {
	cfg, err := utils.MakeConfig(ctx)
	if err != nil {
		println(err.Error())
		return err
	}

	// get client
	chainClient, err := utils.MakeGnfd(cfg)
	if err != nil {
		println(err.Error())
		return err
	}
	spClient := utils.MakeGfSpClient(cfg)

	// check swapIn info
	sp, err := chainClient.QuerySP(ctx.Context, cfg.SpAccount.SpOperatorAddress)
	if err != nil {
		println(err.Error())
		return err
	}
	vgfID := ctx.Uint64(vgfIDFlag.Name)
	swapInInfo, err := chainClient.QuerySwapInInfo(ctx.Context, uint32(vgfID), 0)
	if err != nil {
		println(err.Error())
		return err
	}
	if swapInInfo.GetSuccessorSpId() != sp.GetId() {
		println("sp is not successor sp")
		return errors.New("sp is not successor sp")
	}

	_, executing, err := spClient.QueryRecoverProcess(ctx.Context, uint32(vgfID), uint32(0))
	if err != nil {
		println(err.Error())
		return err
	}
	if executing {
		println("Please wait until the previous recover work is completed")
		return nil
	}

	// trigger
	err = spClient.TriggerRecoverForSuccessorSP(ctx.Context, uint32(vgfID), 0, -1)
	if err != nil {
		println(err.Error())
		return err
	}
	println("Trigger successfully! please waiting process")
	return nil
}

func QueryRecoverProcessAction(ctx *cli.Context) error {
	cfg, err := utils.MakeConfig(ctx)
	if err != nil {
		println(err.Error())
		return err
	}

	gvgID := ctx.Uint64(gvgIDFlag.Name)
	gvgfID := ctx.Uint64(vgfIDFlag.Name)

	if gvgfID != 0 {
		// get client
		chainClient, err := utils.MakeGnfd(cfg)
		if err != nil {
			println(err.Error())
			return err
		}
		familyInfo, err := chainClient.QueryVirtualGroupFamily(ctx.Context, uint32(gvgfID))
		if err != nil {
			println(err.Error())
			return err
		}

		if len(familyInfo.GetGlobalVirtualGroupIds()) == 0 {
			spInfo, err := chainClient.QuerySP(ctx.Context, cfg.SpAccount.SpOperatorAddress)
			if err != nil {
				println(err.Error())
				return err
			}
			if familyInfo.GetPrimarySpId() == spInfo.GetId() {
				println("recover gvg family complete")
				return nil
			} else {
				println("please execute completeSwapIn cmd")
				return nil
			}
		}
	}

	spClient := utils.MakeGfSpClient(cfg)
	gvgstatsList, executing, err := spClient.QueryRecoverProcess(ctx.Context, uint32(gvgfID), uint32(gvgID))
	if err != nil {
		println(err.Error())
		return err
	}

	if executing {
		println("Please Wait, recover executing")
	} else {
		println("recover progress：")
	}

	result := make([]*ProcessResult, 0, len(gvgstatsList))
	for _, gvgStats := range gvgstatsList {
		res := &ProcessResult{
			VirtualGroupId:         gvgStats.VirtualGroupId,
			VirtualGroupFamilyId:   gvgStats.VirtualGroupFamilyId,
			RedundancyIndex:        gvgStats.RedundancyIndex,
			StartAfter:             gvgStats.StartAfter,
			Limit:                  gvgStats.Limit,
			ObjectCount:            gvgStats.ObjectCount,
			FailedObjectTotalCount: gvgStats.FailedObjectTotalCount,
			RecoverFailedObject:    make([]*FailedRecoverObject, 0, len(gvgStats.RecoverFailedObject)),
		}
		switch gvgStats.Status {
		case 0:
			res.Status = ProcessStatus_Recovering
		case 1:
			res.Status = ProcessStatus_Verifying
		case 2:
			res.Status = ProcessStatus_Complete
		}
		for _, obj := range gvgStats.RecoverFailedObject {
			res.RecoverFailedObject = append(res.RecoverFailedObject, &FailedRecoverObject{
				ObjectId:        obj.ObjectId,
				VirtualGroupId:  obj.VirtualGroupId,
				RedundancyIndex: obj.RedundancyIndex,
				RetryTime:       obj.RetryTime,
			})
		}
		result = append(result, res)
	}
	resString, err := json.Marshal(result)
	if err != nil {
		println(err.Error())
		return err
	}
	println(string(resString))
	// create file
	f, err := os.Create("recover_process.json")
	if err != nil {
		println(err.Error())
		return err
	}

	_, err = fmt.Fprintln(f, string(resString))
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func ListGlobalVirtualGroupsBySecondarySPAction(ctx *cli.Context) error {
	cfg, err := utils.MakeConfig(ctx)
	if err != nil {
		println(err.Error())
		return err
	}
	spClient := utils.MakeGfSpClient(cfg)
	spID := ctx.Uint64(targetSPIDFlag.Name)
	res, err := spClient.ListGlobalVirtualGroupsBySecondarySP(ctx.Context, uint32(spID))
	if err != nil {
		println(err.Error())
		return err
	}
	resJson, err := json.Marshal(res)
	if err != nil {
		println(err.Error())
		return err
	}
	println(string(resJson))
	return nil
}

func ListVirtualGroupFamiliesBySpIDAction(ctx *cli.Context) error {
	cfg, err := utils.MakeConfig(ctx)
	if err != nil {
		println(err.Error())
		return err
	}
	spClient := utils.MakeGfSpClient(cfg)
	spID := ctx.Uint64(targetSPIDFlag.Name)
	res, err := spClient.ListVirtualGroupFamiliesSpID(ctx.Context, uint32(spID))
	if err != nil {
		println(err.Error())
		return err
	}
	resJson, err := json.Marshal(res)
	if err != nil {
		println(err.Error())
		return err
	}
	println(string(resJson))
	return nil
}

func CancelSwapInAction(ctx *cli.Context) error {
	cfg, err := utils.MakeConfig(ctx)
	if err != nil {
		println(err.Error())
		return err
	}

	gvgID := ctx.Uint64(gvgIDFlag.Name)
	gvgfID := ctx.Uint64(vgfIDFlag.Name)

	cancelSwapIn := &virtualgrouptypes.MsgCancelSwapIn{
		GlobalVirtualGroupFamilyId: uint32(gvgfID),
		GlobalVirtualGroupId:       uint32(gvgID),
		StorageProvider:            cfg.SpAccount.SpOperatorAddress,
	}

	spClient := utils.MakeGfSpClient(cfg)
	tx, err := spClient.CancelSwapIn(ctx.Context, cancelSwapIn)
	if err != nil {
		println(err.Error())
		return err
	}
	fmt.Printf("tx successfully! tx_hash:%s", tx)
	return nil
}
