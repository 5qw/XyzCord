//+build go1.12

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/5qw/XyzCord/app"
	"github.com/5qw/XyzCord/config"
	"github.com/5qw/XyzCord/logging"
	"github.com/5qw/XyzCord/tview"
	"github.com/5qw/XyzCord/ui/shortcutdialog"
	"github.com/5qw/XyzCord/version"
)

func main() {
	showVersion := flag.Bool("version", false, "Show the version instead of starting XyzCord")
	showShortcutsDialog := flag.Bool("shortcut-dialog", false, "Shows the shortcuts dialog instead of launching XyzCord")
	setConfigDirectory := flag.String("config-dir", "", "Sets the configuration directory")
	setScriptDirectory := flag.String("script-dir", "", "Sets the script directory")
	setConfigFilePath := flag.String("config-file", "", "Sets exact path of the configuration file")
	accountToUse := flag.String("account", "", "Defines which account XyzCord tries to load")
	logPath := flag.String("log", "", "Defines what file we log to")
	flag.Parse()

	if logPath != nil && *logPath != "" {
		logFile, openError := os.OpenFile(*logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if openError != nil {
			panic(openError)
		}
		defer logFile.Close()
		logging.SetDefaultOutput(logFile)
	}

	if setConfigDirectory != nil {
		config.SetConfigDirectory(*setConfigDirectory)
	}
	if setScriptDirectory != nil {
		config.SetScriptDirectory(*setScriptDirectory)
	}
	if setConfigFilePath != nil {
		config.SetConfigFile(*setConfigFilePath)
	}

	//Making sure both the main app and the shortcuts dialog have the
	//correct theme and configuration files.
	configLoadError := config.LoadConfig()
	if configLoadError != nil {
		log.Fatalf("Error loading configuration file (%s).\n", configLoadError.Error())
	}
	themeLoadingError := config.LoadTheme()
	if themeLoadingError != nil {
		panic(themeLoadingError)
	}
	tview.Styles = *config.GetTheme().Theme

	if showShortcutsDialog != nil && *showShortcutsDialog {
		shortcutdialog.RunShortcutsDialogStandalone()
	} else if showVersion != nil && *showVersion {
		fmt.Printf("You are running XyzCord version %s\nKeep in mind that this version might not be correct for manually built versions, as those can contain additional commits.\n", version.Version)
	} else {
		//App that will be reused throughout the process runtime.
		tviewApp := tview.NewApplication()

		if accountToUse != nil && *accountToUse != "" {
			app.SetupApplicationWithAccount(tviewApp, *accountToUse)
		} else {
			app.SetupApplication(tviewApp)
		}

		runError := tviewApp.Run()
		if runError != nil {
			log.Fatalf("Error launching View (%v).\n", runError)
		}
	}
}
