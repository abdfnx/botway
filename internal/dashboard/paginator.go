// forked from: https://github.com/charmbracelet/bubbles/blob/master/paginator/paginator.go
package dashboard

import (
	"github.com/abdfnx/botway/internal/dashboard/components/style"
	tea "github.com/charmbracelet/bubbletea"
)

// Type specifies the way we render pagination.
type Type int

// Pagination rendering options.
const Dots Type = iota

// Paginator is the Bubble Tea model for this user interface.
type Paginator struct {
	Type        Type
	Cursor      int
	Index       int
	Page        int
	PerPage     int
	TotalPages  int
	Width       int
	Height      int
	ActiveDot   string
	InactiveDot string
	Content     string
}

// SetTotalPages is a helper function for calculating the total number of pages
// from a given number of items. It's use is optional since this pager can be
// used for other things beyond navigating sets. Note that it both returns the
// number of total pages and alters the model.
func (m *Paginator) SetTotalPages(items int) int {
	m.Index = items - 1

	if items < 1 {
		return m.TotalPages
	}

	n := items / m.PerPage

	if items%m.PerPage > 0 {
		n++
	}

	m.TotalPages = n

	return n
}

func (m *Paginator) SetHeight(i int) {
	m.Height = i
	m.PerPage = i
}

func (m *Paginator) SetWidth(i int) {
	m.Width = i
}

// ItemsOnPage is a helper function for returning the numer of items on the
// current page given the total numer of items passed as an argument.
func (m Paginator) ItemsOnPage() int {
	if m.Index < 0 {
		return 0
	}

	return 2
}

func (m *Paginator) GetSliceBounds() (start int, end int) {
	length := m.Index + 1
	start = m.Page * m.PerPage

	end = minPaginator(m.Page*m.PerPage+m.PerPage, length)

	return start, end
}

func (m *Paginator) GetSliceStart() int {
	return m.Page * m.PerPage
}

func (m *Paginator) GetCursorIndex() int {
	return m.GetSliceStart() + m.Cursor
}

func (m *Paginator) GoToStart() {
	m.Page = 0
	m.Cursor = 0
}

func (m *Paginator) LineDown() {
	m.Cursor++

	if m.Cursor > m.PerPage-1 || m.GetCursorIndex() > m.Index {
		m.Cursor = 0
	}
}

func (m *Paginator) LineUp() {
	m.Cursor--

	if m.Cursor < 0 {
		if m.Cursor < m.ItemsOnPage() {
			m.Cursor = m.ItemsOnPage() - 1
		} else {
			m.Cursor = m.PerPage - 1
		}
	}
}

// PrevPage is a number function for navigating one page backward. It will not
// page beyond the first page (i.e. page 0).
func (m *Paginator) PrevPage() {
	if m.Page > 0 {
		m.Page--
	} else {
		m.Page = m.TotalPages - 1
	}

	m.Cursor = 0
}

// NextPage is a helper function for navigating one page forward. It will not
// page beyond the last page (i.e. totalPages - 1).
func (m *Paginator) NextPage() {
	if !m.OnLastPage() {
		m.Page++
	} else {
		m.Page = 0
	}

	m.Cursor = 0
}

// OnLastPage returns whether or not we're on the last page.
func (m Paginator) OnLastPage() bool {
	return m.Page == m.TotalPages-1
}

// SetContent set the pager's text content.
func (m *Paginator) SetContent(s string) {
	m.Content = s
}

func (m Paginator) GetContent() string {
	return m.Content
}

func NewPaginator() Paginator {
	return Paginator{
		Page:        0,
		PerPage:     1,
		TotalPages:  1,
		Height:      0,
		Width:       0,
		ActiveDot:   style.PaginatorActiveDot,
		InactiveDot: style.PaginatorInactiveDot,
		Content:     "",
		Type:        Dots,
	}
}

// Update is the Tea update function which binds keystrokes to pagination.
func (m Paginator) Update(msg tea.Msg) (Paginator, tea.Cmd) {
	return m, nil
}

// View renders the pagination to a string.
func (m Paginator) View() string {
	if m.TotalPages <= 1 {
		return ""
	} else {
		return m.dotsView()
	}
}

func (m Paginator) dotsView() string {
	var s string

	for i := 0; i < int(bots_count); i++ {
		if i == m.Page {
			s += m.ActiveDot
			continue
		}

		s += m.InactiveDot
	}

	return s
}

func minPaginator(a, b int) int {
	if a < b {
		return a
	}

	return b
}
