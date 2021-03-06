/**
 * @brief		entry file of program
 * @detail		files sync client (milti-thread)
 * @author		barry
 * @date		2018/4/10
 */
package main

import (
	"./fclient"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	sIP               string // Server IP
	nPort             int    // Server Port
	nTTL              int    // Time To Live (default: 3600 second)
	bDumpLog          bool   // Switch 4 Log Dump
	sLogFile          string // Log File Path
	sAccount          string // Login Name
	sPassword         string // Login Password
	sUncompressFolder string // Folder Which Extract Data
	sProgressFile     string // Progress Status File
	sStopFlagFile     string // Stop Flag File Path
	sDownloadURI      string // URI 4 Download
)

// Package Initialization
func init() {
	/////////////// Parse Arguments From Command Line
	// [Optional]
	flag.StringVar(&sIP, "ip", "127.0.0.1", "file sync server's ip address (default:0.0.0.0)")
	flag.IntVar(&nPort, "port", 31256, "file sync server's listen port (default:31256)")
	flag.IntVar(&nTTL, "ttl", 3600*1, " (time to live (default: 3600 * 1 seconds)")
	flag.StringVar(&sLogFile, "logpath", "./Client.log", "log file's path (default:./Client.log)")
	flag.BoolVar(&bDumpLog, "dumplog", false, "a switch 4 log dump (default:false)")
	flag.StringVar(&sUncompressFolder, "dir", "./FileData/", "data folder path (default :./FileData/)")
	flag.StringVar(&sProgressFile, "progress", "./Progress.xml", "progress file path (default :./Progress.xml)")
	flag.StringVar(&sStopFlagFile, "stopflagfile", "", "stopflag file path (default : NULL)")
	flag.StringVar(&sDownloadURI, "uri", "", "resource's URI 4 download (default : ''; example : SSE/MIN1_TODAY/MIN1_TODAY, SZSE/MIN1_TODAY/MIN1_TODAY)")

	// [Mandatory]
	flag.StringVar(&sAccount, "account", "", "login user name (default: '') ")
	flag.StringVar(&sPassword, "password", "", "login password (default : '') ")
	flag.Parse()
}

// Program Entry Function
func main() {
	/////////////// Set Log File Path
	if true == bDumpLog {
		oLogFile, oLogErr := os.OpenFile(sLogFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if oLogErr != nil {
			log.Fatal("[ERR] main() : a fatal error occur while creating log file ! ", sLogFile)
		}

		log.SetOutput(oLogFile)
	}

	//////////////// Declare && Active FileSync Server / File Scheduler
	log.Println("[INF] [Ver] ######### 1.0.2 ####################")
	log.Println("[INF] [Begin] ##################################")

	for i := 0; i < 6; i++ {
		objSyncClient := &fclient.FileSyncClient{
			DownloadURI:    sDownloadURI,
			StopFlagFile:   sStopFlagFile,
			ServerHost:     fmt.Sprintf("%s:%d", sIP, nPort),
			Account:        sAccount,
			Password:       sPassword,
			TTL:            nTTL,
			ProgressFile:   sProgressFile,
			TotalTaskCount: 1,
			CompleteCount:  0,
		}

		if false == objSyncClient.Initialize() {
			log.Println("[ERR] cannot initialize client obj.")
		} else {
			if true == objSyncClient.DoTasks(sUncompressFolder) {
				break
			}
		}
	}

	log.Println("[INF] [ End ] ##################################")
}
