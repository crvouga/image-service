package homePage

import (
	"imageService/app/admin/adminRoutes"
	"imageService/app/apiDocs/apiDocsRoutes"
	"imageService/app/ctx/appCtx"
	"imageService/app/ctx/reqCtx"
	"imageService/app/home/homeRoutes"
	"imageService/app/projects/projectRoutes"
	"imageService/app/ui/errorPage"
	"imageService/app/ui/mainMenu"
	"imageService/app/ui/page"
	"imageService/app/ui/pageHeader"
	"imageService/app/users/userAccount/userAccountRoutes"
	"imageService/app/users/userAccount/userRole"
	"imageService/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(homeRoutes.HomePage, Respond(ac))
}

const (
	PageTitle = "Home"
)

type Data struct {
	PageHeader pageHeader.PageHeader
	MainMenu   mainMenu.MainMenu
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := reqCtx.FromHttpRequest(ac, r)

		admins, err := ac.UserAccountDB.GetByRole(userRole.Admin)

		if err != nil {
			errorPage.New(err).Redirect(w, r)
			return
		}

		mainMenuData := mainMenu.MainMenu{
			Items: []mainMenu.MainMenuItem{
				{
					Label:       "Projects",
					Description: "Manage your projects",
					URL:         projectRoutes.ToListProjects(),
				},
				{
					Label:       "Account",
					Description: "Manage your account",
					URL:         userAccountRoutes.UserAccountPage,
				},
				{
					Label:       "HTTP API Docs",
					Description: "View the documentation for the HTTP API",
					URL:         apiDocsRoutes.ApiDocsPage,
				},
			},
		}

		if rc.UserAccount.EnsureComputed().IsRoleAdmin {
			mainMenuData.Items = append(mainMenuData.Items, mainMenu.MainMenuItem{
				Label:       "Admin",
				Description: "Manage the admin",
				URL:         adminRoutes.AdminPage,
			})
		} else if len(admins) == 0 {
			mainMenuData.Items = append(mainMenuData.Items, mainMenu.MainMenuItem{
				Label:       "Claim Admin",
				Description: "Claim the admin role",
				URL:         adminRoutes.ClaimAdmin,
			})
		}

		data := Data{
			PageHeader: pageHeader.PageHeader{
				Title: PageTitle,
			},
			MainMenu: mainMenuData,
		}

		page.Respond(data, static.GetSiblingPath("homePage.html"))(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
