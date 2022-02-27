package launcher

import (
	"sort"
	"strings"

	"github.com/invisible-train-40/zsw-lishi-launcher/metrics"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var AppRegistry = map[string]*AppDef{}

func RegisterApp(appDef *AppDef) {
	UserLog.Debug("registering app", zap.String("app_id", appDef.ID))
	AppRegistry[appDef.ID] = appDef
}

func GetMetricAppMeta() map[string]*metrics.AppMeta {
	mapping := make(map[string]*metrics.AppMeta)
	for _, appDef := range AppRegistry {
		mapping[appDef.MetricsID] = &metrics.AppMeta{
			Title: appDef.Title,
			ID:    appDef.ID,
		}
	}
	return mapping
}

var RegisterCommonFlags func(cmd *cobra.Command) error

func RegisterFlags(cmd *cobra.Command) error {
	for _, appDef := range AppRegistry {
		UserLog.Debug("trying to register flags", zap.String("app_id", appDef.ID))
		if appDef.RegisterFlags != nil {
			UserLog.Debug("found non nil flags, registering", zap.String("app_id", appDef.ID))
			err := appDef.RegisterFlags(cmd)
			if err != nil {
				return err
			}
		}
	}
	if RegisterCommonFlags != nil {
		if err := RegisterCommonFlags(cmd); err != nil {
			return err
		}
	}

	return nil
}

func ParseAppsFromArgs(args []string, runByDefault func(string) bool) (apps []string) {
	if len(args) == 0 {
		return ParseAppsFromArgs([]string{"all"}, runByDefault)
	}

	for _, arg := range args {
		chunks := strings.Split(arg, ",")
		for _, app := range chunks {
			app = strings.TrimSpace(app)
			if app == "all" {
				for app := range AppRegistry {
					if !runByDefault(app) {
						continue
					}
					apps = append(apps, app)
				}
			} else {
				if strings.HasPrefix(app, "-") {
					removeApp := app[1:]
					apps = removeElement(apps, removeApp)
				} else {
					apps = append(apps, app)
				}
			}

		}
	}

	sort.Strings(apps)

	return
}

func removeElement(lst []string, el string) (out []string) {
	for _, l := range lst {
		if l != el {
			out = append(out, l)
		}
	}
	return
}
