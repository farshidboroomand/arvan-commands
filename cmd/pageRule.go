package cmd

import (
	"arvan_internal_cli/internals/utils"
	"arvan_internal_cli/pkg/api"
	"arvan_internal_cli/pkg/helpers"
	"arvan_internal_cli/pkg/validator"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/spf13/cobra"
)

const DefaultPageRule = "default"

type RecordList struct {
	Data []PageRule `json:"data"`
}

type PageRule struct {
	ID                 string `json:"id"`
	DomainId           string `json:"domain_id"`
	Seq                int    `json:"seq"`
	UrlType            string `json:"url_type"`
	IsProtected        bool   `json:"is_protected"`
	Url                string `json:"url"`
	CacheLevel         string `json:"cache_level"`
	WafStatus          bool   `json:"waf_status"`
	FwStatus           bool   `json:"fw_status"`
	AccelerationStatus bool   `json:"acceleration_status"`
	Acceleration       string `json:"acceleration"`
	SlinkStatus        bool   `json:"slink_status"`
	Status             bool   `json:"status"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

var pageRuleCmd = &cobra.Command{
	Use:   "pageRule",
	Short: "Interact with Arvan PageRule",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var pageRuleList = &cobra.Command{
	Use:   "list",
	Short: "Get list of a domain pageRules",
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/page-rules",
			Method: "GET",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		responseData, _ := ioutil.ReadAll(res.Body)

		var pageRuleList = new(RecordList)
		_ = json.Unmarshal(responseData, pageRuleList)

		table := utils.NewTable([]string{
			"ID",
			"DomainId",
			"Seq",
			"UrlType",
			"IsProtected",
			"Url",
			"CacheLevel",
			"WafStatus",
			"FwStatus",
			"AccelerationStatus",
			"Acceleration",
			"SlinkStatus",
			"Status",
			"CreatedAt",
			"UpdatedAt",
		})

		for _, foundPageRules := range pageRuleList.Data {
			record := []string{
				foundPageRules.ID,
				foundPageRules.DomainId,
				strconv.Itoa(foundPageRules.Seq),
				foundPageRules.UrlType,
				strconv.FormatBool(foundPageRules.IsProtected),
				foundPageRules.Url,
				foundPageRules.CacheLevel,
				strconv.FormatBool(foundPageRules.WafStatus),
				strconv.FormatBool(foundPageRules.FwStatus),
				strconv.FormatBool(foundPageRules.AccelerationStatus),
				foundPageRules.Acceleration,
				strconv.FormatBool(foundPageRules.SlinkStatus),
				strconv.FormatBool(foundPageRules.Status),
				foundPageRules.CreatedAt,
				foundPageRules.UpdatedAt,
			}
			table.Append(record)
		}

		table.Render()
	},
}

var removeAllPageOfADomain = &cobra.Command{
	Use:   "delete",
	Short: "Remove all pageRule of a specific domain",
	Run: func(cmd *cobra.Command, args []string) {
		_, validationErr := validator.IsDomain(DomainName)

		if validationErr != nil {
			err := helpers.ToBeColored{Expression: validationErr.Error()}
			err.StdoutError().StopExecution()
		}

		request := api.RequestBag{
			URL:    Config.GetUrl() + "/domains/" + DomainName + "/page-rules",
			Method: "GET",
		}

		res, err := request.Do()

		if err != nil {
			err := helpers.ToBeColored{Expression: err.Error()}
			err.StdoutError().StopExecution()
		}

		defer res.Body.Close()

		api.HandleResponseErr(res)

		responseData, _ := ioutil.ReadAll(res.Body)

		var pageRuleList = new(RecordList)
		_ = json.Unmarshal(responseData, pageRuleList)

		var ids = []string{}
		for _, foundPageRules := range pageRuleList.Data {
			if foundPageRules.UrlType != DefaultPageRule {
				ids = append(ids, foundPageRules.ID)
			}
		}

		for _, id := range ids {
			request := api.RequestBag{
				URL:    Config.GetUrl() + "/domains/" + DomainName + "/page-rules/" + id,
				Method: "DELETE",
			}

			res, err := request.Do()

			if err != nil {
				err := helpers.ToBeColored{Expression: err.Error()}
				err.StdoutError().StopExecution()
			}

			api.HandleResponseErr(res)
			fmt.Printf("%s just deleted, may his soul rest in peace :( \n", id)
		}

		defer res.Body.Close()
		fmt.Println("finished :)")
	},
}

func init() {
	rootCmd.AddCommand(pageRuleCmd)
	pageRuleCmd.AddCommand(pageRuleList)
	pageRuleCmd.AddCommand(removeAllPageOfADomain)

	pageRuleList.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	pageRuleList.MarkFlagRequired("name")

	removeAllPageOfADomain.Flags().StringVarP(&DomainName, "name", "n", "", helpDescriptions["domain-name"])
	removeAllPageOfADomain.MarkFlagRequired("name")
}
