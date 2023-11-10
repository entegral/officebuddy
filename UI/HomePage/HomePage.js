import {
  View,
  Text,
  StyleSheet,
  ScrollView,
} from "react-native";
import Calendar from "./Calendar/Calendar";
import CoworkerList from "./CoworkerList/CoworkerList";

export default function HomePage({

}){

  const days =[
    {
      number: 20,
      day: 'Mon',
      slug: 'Monday',
      active: true,
      coworkers: [
        'Kevin Lasher',
        'Elisa Cuan',
        'Robert Bruce',
        'Marquita Nowell'
      ],
    },
    {
      number: 21,
      day: 'Tue',
      slug: 'Tuesday',
      active: true,
      coworkers: [
        'Kevin Lasher',
        'Robert Bruce',
        'Marquita Nowell'
      ],
    },
    {
      number: 22,
      day: 'Wed',
      slug: 'Free Lunch Wednesday',
      active: false,
      coworkers: [
        'Marquita Nowell'
      ],
    },
    {
      number: 23,
      day: 'Thu',
      slug: 'Thursday',
      active: false,
      coworkers: [
        'Kevin Lasher',
        'Marquita Nowell'
      ],
    },
    {
      number: 24,
      day: 'Fri',
      slug: 'Flex Day Friday',
      active: false,
      coworkers: [
        'Kevin Lasher',
        'Robert Bruce',
        'Marquita Nowell'
      ],
    },
    {
      number: 25,
      day: 'Sat',
      slug: 'Saturday',
      active: false,
      coworkers: [
        'Kevin Lasher',
        'Robert Bruce',
        'Marquita Nowell'
      ],
    },
    {
      number: 26,
      day: 'Sun',
      slug: 'Sunday',
      active: false,
      coworkers: [
        'Kevin Lasher',
        'Robert Bruce',
        'Marquita Nowell'
      ],
    },
  ]

  //create an array of coworker objects
  const coworkersByDay = days.map((day) => {
    const obj = {
      day: day.slug,
      coworkers: day.coworkers
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
        />
      </View>

      <ScrollView
        style={styles.coworkerListContainer}
      >
        <CoworkerList
          coworkersByDay={coworkersByDay}
        />
      </ScrollView>
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