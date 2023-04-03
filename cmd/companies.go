package cmd

import (
	"strings"

	"github.com/akecel/gabbro/data"
	"github.com/akecel/gabbro/responses"
	"github.com/akecel/gabbro/utils"

	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/spf13/cobra"
)

var CompaniesCmd = &cobra.Command{
	Use:   "companies [Game name]",
	Short: "Get information about the companies involved in a game",
	Long:  "Get informations about companies involved in a game by using it name for searching in IGDB database.",
	Args:  cobra.MinimumNArgs(1),
	Run:   GetCompanies,
}

func init() {
	RootCmd.AddCommand(CompaniesCmd)
}

func GetCompanies(cmd *cobra.Command, args []string) {
	gamesList, err := data.GetGamesDataByName(strings.Join(args, " "), 1); if err != nil {
		responses.PrintErrorResponse("Game not found")
	}

	game := gamesList[0]

	gameInvolvedCompanies, err := data.GetInvolvedCompaniesDataByIDs(game.InvolvedCompanies, len(game.InvolvedCompanies)); if err != nil {
		responses.PrintErrorResponse("No companies found for this game")
	}

	for i := 0; i < len(gameInvolvedCompanies); i++ {
		company, _ := data.GetCompaniesDataByID(gameInvolvedCompanies[i].Company)

		var companyLogoURL string
		companyLogo, err := data.GetLogosDataByID(company.Logo); if err == nil {
			companyLogoURL = utils.ReconstructImgURL(companyLogo.URL)
		}
		
		response := responses.CompanyResponse{
			Name:           company.Name,
			Description:    company.Description,
			FoundationDate: utils.ParseTimeStampToString(company.StartDate, true),
			URL:            company.URL,
		}

		responses.PrintHeader(GetInvolvedCompaniesRole(gameInvolvedCompanies[i]))

		if len(companyLogoURL) > 0 {
			responses.PrintImageResponse(companyLogoURL)
		}
		
		responses.PrintResponse(response)
	}
}

func GetInvolvedCompaniesRole(company *igdb.InvolvedCompany) string {
	isDeveloper := company.Developer
	isPublisher := company.Publisher
	isPorting := company.Porting
	isSupporting := company.Supporting

	var role string
	if isDeveloper {
		role = "Developer"
	} else if isPublisher {
		role = "Publisher"
	} else if isPorting {
		role = "Porting"
	} else if isSupporting {
		role = "Supporting"
	} else {
		role = "Unknown"
	}

	return role
}
