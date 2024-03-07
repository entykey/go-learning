/*
package main

import (
	"errors"
	"fmt"
	"log"
	"net"
)

// useful links:
// https://stackoverflow.com/questions/27410764/dial-with-a-specific-address-interface-golang
// https://stackoverflow.com/questions/22751035/golang-distinguish-ipv4-ipv6
func GetInterfaceIpv4Addr(interfaceName string) (addr string, err error) {
	var (
		ief      *net.Interface
		addrs    []net.Addr
		ipv4Addr net.IP
	)
	if ief, err = net.InterfaceByName(interfaceName); err != nil { // get interface
		return
	}
	if addrs, err = ief.Addrs(); err != nil { // get addresses
		return
	}
	for _, addr := range addrs { // get ipv4 address
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			break
		}
	}
	if ipv4Addr == nil {
		return "", errors.New(fmt.Sprintf("interface %s don't have an ipv4 address\n", interfaceName))
	}
	return ipv4Addr.String(), nil
}

func main() {
	// Make sure that you're specifying the correct name of
	// the network interface. On different systems, network interfaces
	// might have different names. For example, on Linux, it might be "eth0", "eth1", etc., while on macOS, it could be "en0", "en1", etc.
	interfaceName := "en0" // Change to the interface name you want to get the IPv4 address for
	addr, err := GetInterfaceIpv4Addr(interfaceName)
	if err != nil {
		log.Fatalf("Error getting IPv4 address for interface %s: %v", interfaceName, err)
	}
	fmt.Printf("IPv4 address for interface %s: %s\n", interfaceName, addr)
}
*/

// perform an ARP scan and list the devices on the network
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Error getting network interfaces: %v", err)
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Printf("Error getting addresses for interface %s: %v", iface.Name, err)
			continue
		}

		fmt.Printf("Interface: %s\n", iface.Name)
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			ip := ipNet.IP
			if ip.To4() == nil {
				continue
			}

			fmt.Printf("  IP Address: %s\n", ip.String())
		}

		// Perform ARP scan
		iface, err := net.InterfaceByName(iface.Name)
		if err != nil {
			log.Printf("Error getting interface %s: %v", iface.Name, err)
			continue
		}

		neighbors, err := iface.Addrs()
		if err != nil {
			log.Printf("Error getting neighbors for interface %s: %v", iface.Name, err)
			continue
		}

		fmt.Println("  Devices in network:")
		for _, neighbor := range neighbors {
			ipNet, ok := neighbor.(*net.IPNet)
			if !ok {
				continue
			}

			ip := ipNet.IP
			if ip.To4() == nil {
				continue
			}

			// Exclude broadcast and network address
			if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() || ip.IsMulticast() || ip.IsUnspecified() {
				continue
			}

			fmt.Printf("    %s\n", ip.String())
		}
	}
}
