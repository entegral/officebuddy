import React, { useState, useEffect } from 'react';
import {
  View,
  Text,
  StyleSheet,
  ScrollView,
  TouchableOpacity,
} from "react-native";
import Calendar from "./Calendar/Calendar";
import CoworkerList from "./CoworkerList/CoworkerList";

export default function HomePage({
  user,
  buttonBase,
  buttonText
}){

  //eventually this will be set by the initial api return
  const [initialDays, setInitialDays] = useState({
    20: true,
    21: true,
    22: false,
    23: false,
    24: false,
    25: false,
    26: false,
  });
  const [days, setDays] = useState([
    {
      number: 20,
      day: 'Mon',
      slug: 'Monday',
      coworkers: [
        {name: 'Elisa Cuan', id: 2},
        {name: 'Robert Bruce', id: 3},
        {name: 'Marquita Nowell', id: 4},
      ],
    },
    {
      number: 21,
      day: 'Tue',
      slug: 'Tuesday',
      coworkers: [
        {name: 'Robert Bruce', id: 3},
        {name: 'Marquita Nowell', id: 4},
      ],
    },
    {
      number: 22,
      day: 'Wed',
      slug: 'Free Lunch Wednesday',
      coworkers: [
        {name: 'Marquita Nowell', id: 4},
      ],
    },
    {
      number: 23,
      day: 'Thu',
      slug: 'Thursday',
      coworkers: [
        {name: 'Marquita Nowell', id: 4},
      ],
    },
    {
      number: 24,
      day: 'Fri',
      slug: 'Flex Day Friday',
      coworkers: [
        {name: 'Robert Bruce', id: 3},
        {name: 'Marquita Nowell', id: 4},
      ],
    },
    {
      number: 25,
      day: 'Sat',
      slug: 'Saturday',
      coworkers: [
        {name: 'Elisa Cuan', id: 2},
        {name: 'Robert Bruce', id: 3},
        {name: 'Marquita Nowell', id: 4},
      ],
    },
    {
      number: 26,
      day: 'Sun',
      slug: 'Sunday',
      coworkers: [
        {name: 'Elisa Cuan', id: 2},
        {name: 'Robert Bruce', id: 3},
        {name: 'Marquita Nowell', id: 4},
      ],
    },
  ]);
  const [activeDays, setActiveDays] = useState({
  })
  const [hasChanged, setHasChanged] = useState(false);

  useEffect(() => {
    const nDays = {...initialDays}
    setActiveDays(nDays);
  }, []);

  useEffect(() => {
    const nUser = { ...user, name: `${user.name} (You)`}
    const newDays = days.map((day) => {
      //check if user is in coworkers array by id
      if (
        activeDays[day.number] &&
        !day.coworkers.find((coworker) => coworker.id === user.id)) {
        day.coworkers.unshift(nUser);
      }
      if (
        !activeDays[day.number] &&
        day.coworkers.find((coworker) => coworker.id === user.id)
      ) {
        day.coworkers = day.coworkers.filter((coworker) => coworker.id !== user.id);
      }
      return day;
    })
    setDays(newDays);

  }, [activeDays]);
  
  const activeDayHandler = (day, value) => {
    const nDays = {
      ...activeDays,
      [day]: value,
    }
    const areEqual = Object.keys(nDays).every(key => nDays[key] === initialDays[key]);
    if (!areEqual) {
      setHasChanged(true);
    } else {
      setHasChanged(false);
    }
    setActiveDays(nDays)
  }

  const handleScheduleSubmit = () => {
    //replace with a POST
    setInitialDays({...activeDays});
    setHasChanged(false);
  }


  //create an array of coworker objects
  const coworkersByDay = days.map((day) => {
    const coWorkers = day.coworkers;
    const obj = {
      day: day.slug,
      coworkers: coWorkers
    }
    return obj
  })

  return (
    <View
      style={styles.container}
    >
      <View
        style={{
          flex: 0,
          maxHeight: '33%'
        }}
      >
        <Text
          style={styles.mainText}
        >
          SELECT YOUR AVAILABILITY
        </Text>
        <Calendar
          days={days}
          activeDays={activeDays}
          activeDayHandler={activeDayHandler}
        />
      </View>

      <ScrollView
        style={styles.coworkerListContainer}
      >
        <CoworkerList
          coworkersByDay={coworkersByDay}
        />
      </ScrollView>
      {hasChanged && (
       <View
        style={{
          flex: 0,
          maxHeight: '33%',
        }}>
          <TouchableOpacity
            style={buttonBase}
            onPress={() => handleScheduleSubmit()}
          >
            <Text
              style={buttonText}
            >
              Submit My Plan
            </Text>
          </TouchableOpacity>
        </View> 
      )}
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    paddingTop: 48,
    flex: 1,
  },
  mainText: {
    color: '#fff',
    fontWeight: 700,
  },
  coworkerListContainer: {
    flex: 0,
    maxHeight: '67%',
    overflow: "hidden",
  }
});