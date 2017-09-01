

func main() {
	client := new(ImClient)

	con,err := client.login()
	if err != nil {
		return
	}
	go client.sendmsg(con, 1000 * time.Millisecond)
	client.handleRev(con)
}
