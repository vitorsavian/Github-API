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
	"vitorsavian/github-api/internal/adapters/controllers"
	"vitorsavian/github-api/internal/adapters/rest"
	"vitorsavian/github-api/internal/adapters/services/git/github"
	"vitorsavian/github-api/internal/infrastructure/env"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

//func initTracer() func() {
//	jeagerEndpoint := os.Getenv("JAEGER_ENDPOINT")
//
//	if jeagerEndpoint == "" {
//		jeagerEndpoint = "http://127.0.0.1:14268/api/traces"
//	}
//
//	// Create and install Jaeger export pipeline.
//	flush, err := jaeger.InstallNewPipeline(
//		jaeger.WithCollectorEndpoint(jeagerEndpoint),
//		jaeger.WithProcess(jaeger.Process{
//			ServiceName: serviceName,
//			Tags: []attribute.KeyValue{
//				attribute.String("exporter", "jaeger"),
//			},
//		}),
//		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
//	)
//	if err != nil {
//		level.Error(logger).Log(
//			"msg", "cannot create tracer",
//			"err", err)
//		os.Exit(1)
//	}
//
//	otel.SetTextMapPropagator(
//		propagation.NewCompositeTextMapPropagator(
//			propagation.TraceContext{},
//			propagation.Baggage{}))
//
//	return flush
//}
//
//func initMeter() *prometheus.Exporter {
//	exporter, err := prometheus.InstallNewPipeline(prometheus.Config{})
//	if err != nil {
//		level.Error(logger).Log(
//			"msg", "failed to initialize Prometheus exporter",
//			"err", err)
//		os.Exit(1)
//	}
//
//	meter := global.Meter(meterName)
//
//	ctx := context.Background()
//
//	// Init the metrics
//	reqCounter = metric.Must(meter).NewFloat64Counter(
//		"http_requests_total",
//		metric.WithDescription("Total number of requests"))
//	reqCounter.Add(ctx, float64(0), commonLabels...)
//
//	errCounter = metric.Must(meter).NewFloat64Counter(
//		"http_errors_total",
//		metric.WithDescription("Total number of errors"))
//	errCounter.Add(ctx, float64(0), commonLabels...)
//
//	// Start collecting runtime metrics
//	if err = runtime.Start(); err != nil {
//		level.Error(logger).Log(
//			"msg", "failed to initialize runtime metrics collection",
//			"err", err)
//		os.Exit(1)
//	}
//
//	// Start collecting host metrics
//	if err = host.Start(); err != nil {
//		level.Error(logger).Log(
//			"msg", "failed to initialize host metrics collection",
//			"err", err)
//		os.Exit(1)
//	}
//
//	return exporter
//}

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
		fmt.Println("Server called")
		envs := env.GetEnvironment()

		gitService := github.NewGitService()
		gitController := controllers.NewGitController(gitService, envs)

		apiHandler := rest.NewHandler(*gitController, envs.Port)
		restApi, err := apiHandler.NewApi()
		if err != nil {
			errors.Wrap(err, "failed to initialize rest api")
			return
		}

		restErr := restApi.Run()

		quit := notifyShutdown()
		select {
		case err := <-restErr:
			errors.Wrap(err, "failed while running rest api")
			log.Fatalln(err.Error())
			return
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
