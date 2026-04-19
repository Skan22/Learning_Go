package main

import (
	"github.com/Skan22/LEARNING_GO/DNS_Sniffer/internal/config"
	"github.com/Skan22/LEARNING_GO/DNS_Sniffer/internal/capture"
	"github.com/Skan22/LEARNING_GO/DNS_Sniffer/internal/dns"
	"github.com/Skan22/LEARNING_GO/DNS_Sniffer/internal/kafka"

	"os"
    "os/signal"
    "syscall"
)
