// 	confArgs = "config.toml"
// 	if _, err := toml.DecodeFile(confArgs, &params); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	if params.Jackett.API == "" {
// 		param := os.Getenv("JACKETT")
// 		if param != "" {
// 			params.Jackett.API = param
// 		} else {
// 			log.Fatal("There is no Jackett API configured.")
// 		}
// 	}
//
// 	if len(params.Dest) < 1 {
// 		log.Fatal("There is no Destination configured.")
// 	}
//
// 	for _, dest := range params.Dest {
// 		if dest.API == "" {
// 			param := os.Getenv(strings.ToUpper(dest.Name))
//
// 			if param != "" {
// 				params.GetDestination(dest.Name).API = param
// 			} else {
// 				if len(params.Dest) == 1 {
// 					log.Fatalf("There is no %s API configured.", dest.Name)
// 				} else {
// 					log.Printf("There is no %s API configured.", dest.Name)
// 				}
// 			}
// 		}
//
// 	}
// }

// func main() {

// ParseOptions()
// fmt.Println(config.Params)

// flag.Usage = func() {
// 	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
// }
//
// flag.StringVar(&confArgs, "config", "config.toml", "Config File")

// radarrCommand := flag.NewFlagSet("radarr", flag.ExitOnError)
// radarrPath := radarrCommand.String("url", "", "Url/IP of Radarr")
// radarrApi := radarrCommand.String("api", "", "API of Radarrr")

// sonarrCommand := flag.NewFlagSet("sonarr", flag.ExitOnError)
// sonarrPath := sonarrCommand.String("url", "", "Url/IP of Sonarr")
// sonarrApi := sonarrCommand.String("api", "", "API of Sonarr")
//
// jackettCommand := flag.NewFlagSet("sonarr", flag.ExitOnError)
// jackettPath := jackettCommand.String("url", "", "Url/IP of Jackett")
// jackettApi := jackettCommand.String("api", "", "API of Jackett")
//
// flag.Parse()
//
// if _, err := toml.DecodeFile(confArgs, &params); err != nil {
// 	log.Fatal(err)
// }
//
// if len(os.Args) == 1 {
// 	fmt.Println("usage: jrs <command> [<args>]")
// 	fmt.Println("The commands are: ")
// 	fmt.Println(" radarr   Radarr options")
// 	fmt.Println(" sonarr  Sonarr options")
// 	fmt.Println(" jackett  Jackett options")
// 	return
// }
//
// switch os.Args[1] {
// case "radarr":
// 	radarrCommand.Parse(os.Args[2:])
// case "sonarr":
// 	sonarrCommand.Parse(os.Args[2:])
// case "jackett":
// 	jackettCommand.Parse(os.Args[2:])
// default:
// 	fmt.Printf("%q is not valid command.\n", os.Args[1])
// 	os.Exit(2)
// }
//
// if radarrCommand.Parsed() {
// 	if *radarrPath == "" {
// 		fmt.Println("Please supply the Radarr path using -url option.")
// 		return
// 	} else if *radarrApi == "" {
// 		fmt.Println("Please supply the Sonarr path using -api option.")
// 		return
// 	}
// }
//
// if sonarrCommand.Parsed() {
// 	if *sonarrPath == "" {
// 		fmt.Println("Please supply the Radarr path using -url option.")
// 		return
// 	} else if *sonarrApi == "" {
// 		fmt.Println("Please supply the Sonarr path using -api option.")
// 		return
// 	}
// }
//
// if jackettCommand.Parsed() {
// 	if *jackettPath == "" {
// 		fmt.Println("Please supply the Radarr path using -url option.")
// 		return
// 	} else if *jackettApi == "" {
// 		fmt.Println("Please supply the Sonarr path using -api option.")
// 		return
// 	}
// }

// fmt.Printf("Your message is sent to %q.\n", *recipientFlag)
// fmt.Printf("Message: %q.\n", *messageFlag)

// RadarrDeleteAllIndexers()
// RadarrAddAllIndexes()
// SonarrDeleteAllIndexers()
// SonarrAddAllIndexes()
// SonarrTestAllIndexers()
// a := app.NewApp(params)
// a.AddAllPublicIndexers()

// }*/

package main

import (
	"jrs/cmd/root"
)

func main() {
	root.RootCmd.Execute()
	// config.ParseConfigFile()

	// main := reflect.ValueOf(config.Params)
	// mainElem := main.Elem()
	// items := mainElem.FieldByName("Dest")
	// config.Params.ChangeParams("jackett", "ip", "qwe")

	// config.Params.SaveFile("lala.toml")

}

// func main() {
// 	var echoTimes int
//
// 	var cmdPrint = &cobra.Command{
// 		Use:   "print [string to print]",
// 		Short: "Print anything to the screen",
// 		Long: `print is for printing anything back to the screen.
// For many years people have printed back to the screen.`,
// 		Args: cobra.MinimumNArgs(1),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			fmt.Println("Print: " + strings.Join(args, " "))
// 		},
// 	}
//
// 	var cmdEcho = &cobra.Command{
// 		Use:   "echo [string to echo]",
// 		Short: "Echo anything to the screen",
// 		Long: `echo is for echoing anything back.
// Echo works a lot like print, except it has a child command.`,
// 		Args: cobra.MinimumNArgs(1),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			fmt.Println("Print: " + strings.Join(args, " "))
// 		},
// 	}
//
// 	var cmdTimes = &cobra.Command{
// 		Use:   "times [# times] [string to echo]",
// 		Short: "Echo anything to the screen more times",
// 		Long: `echo things multiple times back to the user by providing
// a count and a string.`,
// 		Args: cobra.MinimumNArgs(1),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			for i := 0; i < echoTimes; i++ {
// 				fmt.Println("Echo: " + strings.Join(args, " "))
// 			}
// 		},
// 	}
//
// 	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
//
// 	var rootCmd = &cobra.Command{Use: "app"}
// 	rootCmd.AddCommand(cmdPrint, cmdEcho)
// 	cmdEcho.AddCommand(cmdTimes)
// 	rootCmd.Execute()
// }
