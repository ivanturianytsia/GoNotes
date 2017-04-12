package notes

import (
	"time"
  "os"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
  s = getSession()
)

// Profile - is the memory representation of one user profile
type Note struct {
	Title string `json: "title"`
	Content string `json: "content"`
	Created time.Time `json: "username"`
	Updates time.Time `json: "username"`
}

func getSession() *mgo.Session {
  session, err := mgo.Dial(os.Getenv("DB_URL"))
  if err != nil {
		panic(err)
    return nil
	}
	session.SetMode(mgo.Monotonic, true)
  return session
}

func GetNotes() (*[]Note, error) {
	session := s.Copy()
	defer session.Close()
	c := session.DB("NotesApp").C("Notes")

	notes := make([]Note, 0, 10)
	err := c.Find(bson.M{}).All(&notes)

	return &notes, err
}

// func (p *Profile) CreateOrUpdateProfile() bool {
//   session := s.Copy()
// 	defer session.Close()
// 	c := session.DB("NotesApp").C("Notes")
//
//
// 	_ , err = c.UpsertId( p.Name, p )
// 	if err != nil {
// 		log.Println("Error creating Profile: ", err.Error())
// 		return false
// 	}
// 	return true
// }


// // ShowProfile - Returns the profile in the Profiles Collection with name equal to the id parameter (id == name)
// func ShowProfile(id string) Profile {
// 	session, err := mgo.Dial("localhost:27017")
// 	if err != nil {
// 		log.Println("Could not connect to mongo: ", err.Error())
// 		return Profile{}
// 	}
// 	defer session.Close()
//
// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)
//
// 	c := session.DB("ProfileService").C("Profiles")
// 	profile := Profile{}
// 	err = c.Find(bson.M{"name": id}).One(&profile)
//
// 	return profile
// }
//
//
// // DeleteProfile - Deletes the profile in the Profiles Collection with name equal to the id parameter (id == name)
// func DeleteProfile(id string) bool {
// 	session, err := mgo.Dial("localhost:27017")
// 	if err != nil {
// 		log.Println("Could not connect to mongo: ", err.Error())
// 		return false
// 	}
// 	defer session.Close()
//
// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)
//
// 	c := session.DB("ProfileService").C("Profiles")
// 	err = c.RemoveId(id)
//
// 	return true
// }
//
//
// // CreateOrUpdateProfile - Creates or Updates (Upsert) the profile in the Profiles Collection with id parameter
