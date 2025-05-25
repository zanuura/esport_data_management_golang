package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	green = "\033[32m"
	reset = "\033[0m"
)

type Team struct {
	Name         string
	Matches      int
	Wins         int
	Losses       int
	ScoreFor     int
	ScoreAgainst int
}

var teams []Team

type Match struct {
	TeamA  string
	TeamB  string
	ScoreA int
	ScoreB int
}

var matchHistory []Match

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(green + "\n--- Aplikasi Pengelolaan E-Sports ---" + reset)
		fmt.Println("1. Tambah Tim")
		fmt.Println("2. Tambah Hasil Pertandingan")
		fmt.Println("3. Tampilkan Klasemen")
		fmt.Println("4. Cari Tim (Sequential/Binary Search)")
		fmt.Println("5. Urutkan Tim (Selection / Insertion Sort)")
		fmt.Println("6. Statistik Tim Terbaik")
		fmt.Println("7. Lihat Semua Tim")
		fmt.Println("8. Lihat Semua Pertandingan")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		switch scanner.Text() {
		case "1":
			addTeam(scanner)
		case "2":
			addMatchResult(scanner)
		case "3":
			showStandings()
		case "4":
			searchTeam(scanner)
		case "5":
			sortTeams(scanner)
		case "6":
			showBestTeam()
		case "7":
			viewAndEditTeams(scanner)
		case "8":
			showAllMatches()
		case "9":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func addTeam(scanner *bufio.Scanner) {
	fmt.Println("Masukkan detail tim:")

	fmt.Print("Nama tim: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Jumlah pertandingan: ")
	scanner.Scan()
	matches, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Jumlah kemenangan: ")
	scanner.Scan()
	wins, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Jumlah kekalahan: ")
	scanner.Scan()
	losses, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Skor yang dicetak: ")
	scanner.Scan()
	scoreFor, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Skor yang diterima: ")
	scanner.Scan()
	scoreAgainst, _ := strconv.Atoi(scanner.Text())

	teams = append(teams, Team{
		Name:         name,
		Matches:      matches,
		Wins:         wins,
		Losses:       losses,
		ScoreFor:     scoreFor,
		ScoreAgainst: scoreAgainst,
	})

	fmt.Println("✅ Tim berhasil ditambahkan.")
}

func addMatchResult(scanner *bufio.Scanner) {
	fmt.Print("Nama tim A: ")
	scanner.Scan()
	a := scanner.Text()

	fmt.Print("Nama tim B: ")
	scanner.Scan()
	b := scanner.Text()

	fmt.Print("Skor tim A: ")
	scanner.Scan()
	scoreA, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Skor tim B: ")
	scanner.Scan()
	scoreB, _ := strconv.Atoi(scanner.Text())

	updateTeamResult(a, scoreA, scoreB)
	updateTeamResult(b, scoreB, scoreA)
	matchHistory = append(matchHistory, Match{
		TeamA:  a,
		TeamB:  b,
		ScoreA: scoreA,
		ScoreB: scoreB,
	})
	fmt.Println("Hasil pertandingan dicatat.")
}

func updateTeamResult(name string, scoreFor, scoreAgainst int) {
	for i := range teams {
		if teams[i].Name == name {
			teams[i].Matches++
			teams[i].ScoreFor += scoreFor
			teams[i].ScoreAgainst += scoreAgainst
			if scoreFor > scoreAgainst {
				teams[i].Wins++
			} else {
				teams[i].Losses++
			}
			return
		}
	}
}

func showAllMatches() {
	if len(matchHistory) == 0 {
		fmt.Println("❌ Belum ada hasil pertandingan.")
		return
	}

	fmt.Println(green + "\nDaftar Hasil Pertandingan:")
	fmt.Println("No | Tim A              | Skor A | VS | Skor B | Tim B")
	fmt.Println("--------------------------------------------------------" + reset)
	for i, match := range matchHistory {
		fmt.Printf("%-2d | %-18s | %-6d | VS | %-6d | %s\n",
			i+1, match.TeamA, match.ScoreA, match.ScoreB, match.TeamB)
	}
}

func showStandings() {
	if len(teams) == 0 {
		fmt.Println("Belum ada data tim untuk ditampilkan.")
		return
	}

	sort.SliceStable(teams, func(i, j int) bool {
		if teams[i].Wins == teams[j].Wins {
			return (teams[i].ScoreFor - teams[i].ScoreAgainst) > (teams[j].ScoreFor - teams[j].ScoreAgainst)
		}
		return teams[i].Wins > teams[j].Wins
	})

	fmt.Println(green + "\nKlasemen Tim:")
	fmt.Println("No | Nama Tim           | Main | Menang | Kalah | Skor For | Skor Against")
	fmt.Println("--------------------------------------------------------------------------" + reset)
	for i, t := range teams {
		fmt.Printf("%-2d | %-18s | %-4d | %-6d | %-5d | %-8d | %-13d\n",
			i+1, t.Name, t.Matches, t.Wins, t.Losses, t.ScoreFor, t.ScoreAgainst)
	}
}

func searchTeam(scanner *bufio.Scanner) {
	fmt.Print("Cari nama tim: ")
	scanner.Scan()
	query := strings.ToLower(scanner.Text())

	var results []Team
	for _, t := range teams {
		nameLower := strings.ToLower(t.Name)
		if strings.Contains(nameLower, query) {
			results = append(results, t)
		}
	}

	if len(results) == 0 {
		fmt.Println("❌ Tidak ditemukan tim yang cocok.")
		return
	}

	fmt.Println(green + "\nHasil Pencarian Tim:")
	fmt.Println("No | Nama Tim           | Main | Menang | Kalah | Skor For | Skor Against")
	fmt.Println("--------------------------------------------------------------------------" + reset)
	for i, t := range results {
		fmt.Printf("%-2d | %-18s | %-4d | %-6d | %-5d | %-8d | %-13d\n",
			i+1, t.Name, t.Matches, t.Wins, t.Losses, t.ScoreFor, t.ScoreAgainst)
	}
}

func sortTeams(scanner *bufio.Scanner) {
	if len(teams) == 0 {
		fmt.Println("Belum ada tim untuk diurutkan.")
		return
	}

	fmt.Print("Gunakan Selection Sort atau Insertion Sort? (s/i): ")
	scanner.Scan()
	method := strings.ToLower(scanner.Text())

	var compare func(a, b Team) bool

	if method == "s" {
		fmt.Print("Urut berdasarkan (w: menang, l: kalah, m: match, s: skor): ")
		scanner.Scan()
		criteria := strings.ToLower(scanner.Text())

		compare = getCompareFunc(criteria)

		selectionSort(compare)
		fmt.Println("✅ Diurutkan dengan Selection Sort.")
	} else {
		// Gunakan default sorting berdasarkan kemenangan
		compare = func(a, b Team) bool {
			return a.Wins > b.Wins
		}
		insertionSort(compare)
		fmt.Println("✅ Diurutkan dengan Insertion Sort (default: kemenangan).")
	}

	// Tampilkan hasil urutan
	fmt.Println(green + "\nHasil Pengurutan:")
	fmt.Println("No | Nama Tim           | Main | Menang | Kalah | Skor For | Skor Against")
	fmt.Println("--------------------------------------------------------------------------" + reset)
	for i, t := range teams {
		fmt.Printf("%-2d | %-18s | %-4d | %-6d | %-5d | %-8d | %-13d\n",
			i+1, t.Name, t.Matches, t.Wins, t.Losses, t.ScoreFor, t.ScoreAgainst)
	}
}

func selectionSort(compare func(a, b Team) bool) {
	for i := 0; i < len(teams); i++ {
		maxIdx := i
		for j := i + 1; j < len(teams); j++ {
			if compare(teams[j], teams[maxIdx]) {
				maxIdx = j
			}
		}
		teams[i], teams[maxIdx] = teams[maxIdx], teams[i]
	}
}

func getCompareFunc(criteria string) func(a, b Team) bool {
	switch criteria {
	case "l":
		return func(a, b Team) bool {
			return a.Losses > b.Losses
		}
	case "m":
		return func(a, b Team) bool {
			return a.Matches > b.Matches
		}
	case "s":
		return func(a, b Team) bool {
			return (a.ScoreFor - a.ScoreAgainst) > (b.ScoreFor - b.ScoreAgainst)
		}
	default: // default ke kemenangan
		return func(a, b Team) bool {
			return a.Wins > b.Wins
		}
	}
}

func insertionSort(compare func(a, b Team) bool) {
	for i := 1; i < len(teams); i++ {
		key := teams[i]
		j := i - 1
		for j >= 0 && compare(key, teams[j]) {
			teams[j+1] = teams[j]
			j--
		}
		teams[j+1] = key
	}
}

func showBestTeam() {
	if len(teams) == 0 {
		fmt.Println("Belum ada data tim.")
		return
	}

	// Buat slice baru dengan winrate
	type TeamStat struct {
		Team
		WinRate float64
	}
	var stats []TeamStat

	for _, t := range teams {
		winRate := 0.0
		if t.Matches > 0 {
			winRate = float64(t.Wins) / float64(t.Matches) * 100
		}
		stats = append(stats, TeamStat{t, winRate})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].WinRate > stats[j].WinRate
	})

	fmt.Println(green + "\nStatistik Tim dengan Performa Terbaik (berdasarkan Win Rate):")
	fmt.Println("No | Nama Tim           | Main | Menang | Win Rate (%)")
	fmt.Println("-------------------------------------------------------" + reset)
	for i, stat := range stats {
		fmt.Printf("%-2d | %-18s | %-4d | %-6d | %.2f%%\n",
			i+1, stat.Name, stat.Matches, stat.Wins, stat.WinRate)
		if i == 2 { // hanya tampilkan top 3
			break
		}
	}
}

func viewAndEditTeams(scanner *bufio.Scanner) {
	if len(teams) == 0 {
		fmt.Println("Belum ada tim yang tercatat.")
		return
	}

	fmt.Println("\nDaftar Tim:")
	for i, t := range teams {
		fmt.Printf("%d. %s | Main: %d | Menang: %d | Kalah: %d | Skor: %d:%d\n",
			i+1, t.Name, t.Matches, t.Wins, t.Losses, t.ScoreFor, t.ScoreAgainst)
	}

	fmt.Print("Masukkan nomor tim yang ingin diedit (atau 0 untuk batal): ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	if choice < 1 || choice > len(teams) {
		fmt.Println("❌ Tidak ada tim yang dipilih.")
		return
	}

	team := &teams[choice-1]

	fmt.Printf("Edit data untuk tim '%s'\n", team.Name)
	fmt.Print("Nama baru (enter jika tidak diubah): ")
	scanner.Scan()
	newName := scanner.Text()
	if newName != "" {
		team.Name = newName
	}

	fmt.Print("Jumlah pertandingan: ")
	scanner.Scan()
	team.Matches, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Jumlah kemenangan: ")
	scanner.Scan()
	team.Wins, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Jumlah kekalahan: ")
	scanner.Scan()
	team.Losses, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Skor yang dicetak: ")
	scanner.Scan()
	team.ScoreFor, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Skor yang diterima: ")
	scanner.Scan()
	team.ScoreAgainst, _ = strconv.Atoi(scanner.Text())

	fmt.Println("✅ Tim berhasil diperbarui.")
}
