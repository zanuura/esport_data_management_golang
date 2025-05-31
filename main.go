package main

import (
	"fmt"
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

type TeamStat struct {
	Team
	WinRate float64
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

	// INFINITE LOOP UNTUK MENAMPILKAN MENU
	for {
		fmt.Println(green + "\n--- Aplikasi Pengelolaan E-Sports ---" + reset)
		fmt.Println("1. Tambah Tim")
		fmt.Println("2. Tambah Hasil Pertandingan")
		fmt.Println("3. Tampilkan Klasemen")
		fmt.Println("4. Cari Tim")
		fmt.Println("5. Urutkan Tim (Selection / Insertion Sort)")
		fmt.Println("6. Statistik Tim Terbaik")
		fmt.Println("7. Lihat Semua Tim")
		fmt.Println("8. Lihat Semua Pertandingan")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")

		var input string
		fmt.Scanln(&input)

		switch input {
		case "1":
			addTeam()
		case "2":
			addMatchResult()
		case "3":
			showStandings()
		case "4":
			searchTeam()
		case "5":
			sortTeams()
		case "6":
			showBestTeam()
		case "7":
			viewAndEditTeams()
		case "8":
			showAllMatches()
		case "9":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func addTeam() {
	var name string
	var matches, wins, losses, scoreFor, scoreAgainst int

	fmt.Println("Masukkan detail tim:")
	fmt.Print("Nama tim: ")
	fmt.Scanln(&name)

	fmt.Print("Jumlah pertandingan: ")
	fmt.Scanln(&matches)

	fmt.Print("Jumlah kemenangan: ")
	fmt.Scanln(&wins)

	fmt.Print("Jumlah kekalahan: ")
	fmt.Scanln(&losses)

	fmt.Print("Skor yang dicetak: ")
	fmt.Scanln(&scoreFor)

	fmt.Print("Skor yang diterima: ")
	fmt.Scanln(&scoreAgainst)

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

func addMatchResult() {
	var a, b string
	var scoreA, scoreB int

	fmt.Print("Nama tim A: ")
	fmt.Scanln(&a)

	fmt.Print("Nama tim B: ")
	fmt.Scanln(&b)

	fmt.Print("Skor tim A: ")
	fmt.Scanln(&scoreA)

	fmt.Print("Skor tim B: ")
	fmt.Scanln(&scoreB)

	updateTeamResult(a, scoreA, scoreB)
	updateTeamResult(b, scoreB, scoreA)

	matchHistory = append(matchHistory, Match{
		TeamA:  a,
		TeamB:  b,
		ScoreA: scoreA,
		ScoreB: scoreB,
	})

	fmt.Println("✅ Hasil pertandingan dicatat.")
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

	customSortByWinsAndGoalDiff(teams)

	fmt.Println(green + "\nKlasemen Tim:")
	fmt.Println("No | Nama Tim           | Main | Menang | Kalah | Skor For | Skor Against")
	fmt.Println("--------------------------------------------------------------------------" + reset)
	for i, t := range teams {
		fmt.Printf("%-2d | %-18s | %-4d | %-6d | %-5d | %-8d | %-13d\n",
			i+1, t.Name, t.Matches, t.Wins, t.Losses, t.ScoreFor, t.ScoreAgainst)
	}
}

func searchTeam() {
	var query string
	fmt.Print("Cari nama tim: ")
	fmt.Scanln(&query)

	query = strings.ToLower(query)
	var results []Team
	for _, t := range teams {
		if strings.Contains(strings.ToLower(t.Name), query) {
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

func sortTeams() {
	if len(teams) == 0 {
		fmt.Println("Belum ada tim untuk diurutkan.")
		return
	}

	var method, criteria string
	fmt.Print("Gunakan Selection Sort atau Insertion Sort? (s/i): ")
	fmt.Scanln(&method)

	var compare func(a, b Team) bool

	if method == "s" {
		fmt.Print("Urut berdasarkan (w: menang, l: kalah, m: match, s: skor): ")
		fmt.Scanln(&criteria)
		compare = getCompareFunc(criteria)
		selectionSort(compare)
		fmt.Println("✅ Diurutkan dengan Selection Sort.")
	} else {
		compare = func(a, b Team) bool {
			return a.Wins > b.Wins
		}
		insertionSort(compare)
		fmt.Println("✅ Diurutkan dengan Insertion Sort (default: kemenangan).")
	}

	fmt.Println(green + "\nHasil Pengurutan:")
	fmt.Println("No | Nama Tim           | Main | Menang | Kalah | Skor For | Skor Against")
	fmt.Println("--------------------------------------------------------------------------" + reset)
	for i, t := range teams {
		fmt.Printf("%-2d | %-18s | %-4d | %-6d | %-5d | %-8d | %-13d\n",
			i+1, t.Name, t.Matches, t.Wins, t.Losses, t.ScoreFor, t.ScoreAgainst)
	}
}

func selectionSort(compare func(a, b Team) bool) {
	// FOR LOOP KLASIK
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

func insertionSort(compare func(a, b Team) bool) {
	for i := 1; i < len(teams); i++ {
		key := teams[i]
		j := i - 1
		// WHILE LOOP UNTUK MENGGESER TIM YANG LEBIH RENDAH
		for j >= 0 && compare(key, teams[j]) {
			teams[j+1] = teams[j]
			j--
		}
		teams[j+1] = key
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
	default:
		return func(a, b Team) bool {
			return a.Wins > b.Wins
		}
	}
}

func showBestTeam() {
	if len(teams) == 0 {
		fmt.Println("Belum ada data tim.")
		return
	}

	// type TeamStat struct {
	// 	Team
	// 	WinRate float64
	// }
	var stats []TeamStat

	// SEQUENTIAL SEARCH dengan FOR-RANGE
	for _, t := range teams {
		winRate := 0.0
		if t.Matches > 0 {
			winRate = float64(t.Wins) / float64(t.Matches) * 100
		}
		stats = append(stats, TeamStat{t, winRate})
	}

	sortTeamStatsByWinRate(stats)

	fmt.Println(green + "\nStatistik Tim Terbaik:")
	fmt.Println("No | Nama Tim           | Main | Menang | Win Rate (%)")
	fmt.Println("-------------------------------------------------------" + reset)
	for i, stat := range stats {
		fmt.Printf("%-2d | %-18s | %-4d | %-6d | %.2f%%\n",
			i+1, stat.Name, stat.Matches, stat.Wins, stat.WinRate)
		if i == 2 {
			break
		}
	}
}

func viewAndEditTeams() {
	if len(teams) == 0 {
		fmt.Println("Belum ada tim yang tercatat.")
		return
	}

	fmt.Println("\nDaftar Tim:")
	for i, t := range teams {
		fmt.Printf("%d. %s | Main: %d | Menang: %d | Kalah: %d | Skor: %d:%d\n",
			i+1, t.Name, t.Matches, t.Wins, t.Losses, t.ScoreFor, t.ScoreAgainst)
	}

	var choice int
	fmt.Print("Masukkan nomor tim yang ingin diedit (atau 0 untuk batal): ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(teams) {
		fmt.Println("❌ Tidak ada tim yang dipilih.")
		return
	}

	team := &teams[choice-1]

	fmt.Printf("Edit data untuk tim '%s'\n", team.Name)

	var newName string
	fmt.Print("Nama baru (kosong jika tidak diubah): ")
	fmt.Scanln(&newName)
	if newName != "" {
		team.Name = newName
	}

	fmt.Print("Jumlah pertandingan: ")
	fmt.Scanln(&team.Matches)
	fmt.Print("Jumlah kemenangan: ")
	fmt.Scanln(&team.Wins)
	fmt.Print("Jumlah kekalahan: ")
	fmt.Scanln(&team.Losses)
	fmt.Print("Skor yang dicetak: ")
	fmt.Scanln(&team.ScoreFor)
	fmt.Print("Skor yang diterima: ")
	fmt.Scanln(&team.ScoreAgainst)

	fmt.Println("✅ Tim berhasil diperbarui.")
}

func customSortByWinsAndGoalDiff(arr []Team) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			diffI := arr[i].ScoreFor - arr[i].ScoreAgainst
			diffJ := arr[j].ScoreFor - arr[j].ScoreAgainst

			if arr[j].Wins > arr[i].Wins || (arr[j].Wins == arr[i].Wins && diffJ > diffI) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func sortTeamStatsByWinRate(stats []TeamStat) {
	for i := 0; i < len(stats); i++ {
		for j := i + 1; j < len(stats); j++ {
			if stats[j].WinRate > stats[i].WinRate {
				stats[i], stats[j] = stats[j], stats[i]
			}
		}
	}
}
