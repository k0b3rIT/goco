GOCO 🤖 
=========

GOCOllections framework

![GitHub tag (with filter)](https://img.shields.io/github/v/tag/k0b3rIT/goco)
[![Repo stars](https://img.shields.io/github/stars/k0b3rIT/goco?style=social)](https://github.com/k0b3rIT/goco)
[![Repo forks](https://img.shields.io/github/forks/k0b3rIT/goco?style=social)](https://github.com/k0b3rIT/goco)

Go-collections framework is a set of packages that extend the Go standard library to provide a more complete, convenient and powerful set of data structures.
(Go list, Go set, Go map, Go collections) (Java collections for Go)

## List

```go
//Create a new list from Car objects
carList := NewList[Car](Car{"bmw", "x5"}, Car{"audi", "a4"}, Car{"mercedes", "c180"})

//Add a new car to the list
carList.Add(Car{"bmw", "x6"})

//Remove a car from the list
carList.Remove(Car{"bmw", "x5"})

//Get the first car from the list
car := carList.Get(0)

//Get the size of the list
size := carList.Size()

//Is the list empty?
isEmpty := carList.IsEmpty()

//Convert to map
carMap := ListToMap[Car](l, func(c Car) string { return c.Brand })

//Go back to slice
carSlice := carList.ToSlice()
```



## Set

```go
//Create a new set from Car objects
carSet := NewSet[Car](Car{"bmw", "x5"}, Car{"audi", "a4"}, Car{"mercedes", "c180"})

//Add a new car to the set (existing cars will be overwritten)
carSet.Add(Car{"bmw", "x6"})

//Remove a car from the set
carSet.Remove(Car{"bmw", "x5"})
```

## Map

```go
//Create a new map from Car objects
carMap := NewMap[Car]()

//Add a new car to the map (existing cars will be overwritten)
carMap.Put("bmw", Car{"bmw", "x5"})

//Is the map contains the key?
contains := carMap.ContainsKey("bmw")

//Get the car from the map
car := carMap.Get("bmw")

//Get the keys
keys := carMap.Keys()
```