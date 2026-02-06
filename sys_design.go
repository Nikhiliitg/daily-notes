package dailynotes

// Why you should not use a relational , Nosql and even a Linked list in Google sheet design 

// n an RDBMS, updates are: 1.Row-oriented 2.Transactional 3.Discrete and deliberate

//

type Cell struct {
    ID           string
    Value        interface{}
    Formula      string
    Dependencies map[string]bool
    Dependents   map[string]bool
}

type FormulaEngine struct {
	CellStore *CellStore
}
type CellStore struct {
    Cells map[string]*Cell
}

func update_cell(cell *Cell, newValue interface{}) {
	cell.Value = newValue
}

func recompute_dependents(cell *Cell, engine *FormulaEngine) {
	for dependentID := range cell.Dependents {
		dependentCell := engine.CellStore.GetOrCreate(dependentID)
		// Recompute the dependent cell's value based on its formula
		// This is a placeholder for actual formula evaluation logic
		dependentCell.Value = evaluate_formula(dependentCell.Formula, engine)
		recompute_dependents(dependentCell, engine)
	}
}

func evaluate_formula(formula string, engine *FormulaEngine) interface{} {
	// Placeholder for formula evaluation logic
	return nil
}
type Sheet struct {
    Cells         map[string]*Cell
    RowIndex      []int
    ColumnIndex   []string
    FormulaEngine *FormulaEngine
}

func (cs *CellStore) GetOrCreate(id string) *Cell {
    if cs.Cells[id] == nil {
        cs.Cells[id] = &Cell{ID: id}
    }
    return cs.Cells[id]
}
