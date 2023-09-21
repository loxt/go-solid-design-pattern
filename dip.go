package main

// Dependency Inversion Principle

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

type Research struct {
	relationship Relationships
}

func (r *Research) Investigate() {
	relations := r.relationship.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			println("John has a child called", rel.to.name)
		}
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)

	// high-level module
	research := Research{relationships}
	research.Investigate()
}
