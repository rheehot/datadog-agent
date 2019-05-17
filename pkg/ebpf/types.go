package ebpf

import "fmt"

// KProbeName stores the name of the kernel probes setup for tracing
type KProbeName string

const (
	// TCPv4Connect traces the v4 connect() system call
	TCPv4Connect KProbeName = "kprobe/tcp_v4_connect"
	// TCPv4ConnectReturn traces the return value for the v4 connect() system call
	TCPv4ConnectReturn KProbeName = "kretprobe/tcp_v4_connect"
	// TCPv4DestroySock traces the tcp_v4_destroy_sock system call (called for both ipv4 and ipv6)
	TCPv4DestroySock KProbeName = "kprobe/tcp_v4_destroy_sock"

	// TCPv6Connect traces the v6 connect() system call
	TCPv6Connect KProbeName = "kprobe/tcp_v6_connect"
	// TCPv6ConnectReturn traces the return value for the v6 connect() system call
	TCPv6ConnectReturn KProbeName = "kretprobe/tcp_v6_connect"

	// TCPSendMsg traces the tcp_sendmsg() system call
	TCPSendMsg KProbeName = "kprobe/tcp_sendmsg"
	// TCPCleanupRBuf traces the tcp_cleanup_rbuf() system call
	TCPCleanupRBuf KProbeName = "kprobe/tcp_cleanup_rbuf"
	// TCPClose traces the tcp_close() system call
	TCPClose KProbeName = "kprobe/tcp_close"

	// UDPSendMsg traces the udp_sendmsg() system call
	UDPSendMsg KProbeName = "kprobe/udp_sendmsg"
	// UDPRecvMsg traces the udp_recvmsg() system call
	UDPRecvMsg KProbeName = "kprobe/udp_recvmsg"
	// UDPRecvMsgReturn traces the return value for the udp_recvmsg() system call
	UDPRecvMsgReturn KProbeName = "kretprobe/udp_recvmsg"

	// TCPRetransmit traces the return value for the tcp_retransmit_skb() system call
	TCPRetransmit KProbeName = "kprobe/tcp_retransmit_skb"

	// InetCskAcceptReturn traces the return value for the inet_csk_accept syscall
	InetCskAcceptReturn KProbeName = "kretprobe/inet_csk_accept"
)

// bpfMapName stores the name of the BPF maps storing statistics and other info
type bpfMapName string

const (
	connMap            bpfMapName = "conn_stats"
	tcpStatsMap        bpfMapName = "tcp_stats"
	tcpCloseEventMap   bpfMapName = "tcp_close_events"
	latestTimestampMap bpfMapName = "latest_ts"
	tracerStatusMap    bpfMapName = "tracer_status"
	portBindingsMap    bpfMapName = "port_bindings"
)

// sectionName returns the sectionName for the given BPF map
func (b bpfMapName) sectionName() string {
	return fmt.Sprintf("maps/%s", b)
}