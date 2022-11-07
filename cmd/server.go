/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		// ctx := context.Background()

		// exporter, err := prometheus.New()
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }

		// provider := metric.NewMeterProvider(metric.WithReader(exporter))
		// meter := provider.Meter("github.com/open-telemetry/opentelemetry-go/example/prometheus")

		// go serveMetrics()

		// attrs := []attribute.KeyValue{
		// 	attribute.Key("A").String("B"),
		// 	attribute.Key("C").String("D"),
		// 	attribute.Key("F").String("Y"),
		// }

		// // This is the equivalent of prometheus.NewCounterVec
		// counter, err := meter.SyncFloat64().Counter("foo", instrument.WithDescription("a simple counter"))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// counter.Add(ctx, 5, attrs...)

		// gauge, err := meter.SyncFloat64().UpDownCounter("bar", instrument.WithDescription("a fun little gauge"))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// gauge.Add(ctx, 100, attrs...)
		// gauge.Add(ctx, -25, attrs...)

		// // This is the equivalent of prometheus.NewHistogramVec
		// histogram, err := meter.SyncFloat64().Histogram("baz", instrument.WithDescription("a very nice histogram"))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// histogram.Record(ctx, 23, attrs...)
		// histogram.Record(ctx, 7, attrs...)
		// histogram.Record(ctx, 101, attrs...)
		// histogram.Record(ctx, 105, attrs...)

		// ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
		// <-ctx.Done()

		quit := notifyShutdown()
		select {
		case <-quit:
			log.Println("gracefully shutdown")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func notifyShutdown() chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	return quit
}

// func serveMetrics() {
// 	log.Printf("serving metrics at localhost:2223/metrics")
// 	http.Handle("/metrics", promhttp.Handler())
// 	err := http.ListenAndServe(":2223", nil)
// 	if err != nil {
// 		fmt.Printf("error serving http: %v", err)
// 		return
// 	}
// }
