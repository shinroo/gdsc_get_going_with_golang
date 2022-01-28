package genx

type PersonList []*Person

type PersonToBool func(*Person) bool

func (al PersonList)Filter(f PersonToBool) PersonList {
   var ret PersonList
   for _, a := range al {
      if f(a) {
         ret = append(ret, a)
      }
   }
   return ret
}
