import { View, Text, StyleSheet } from 'react-native';
import CalendarDay from './CalendarDay';

export default function Calendar({}) {
  const days =[
    {
      number: 20,
      day: 'Mon',
      active: true,
      coworkers: 4,
    },
    {
      number: 21,
      day: 'Tue',
      active: true,
      coworkers: 3,
    },
    {
      number: 22,
      day: 'Wed',
      active: false,
      coworkers: 1,
    },
    {
      number: 23,
      day: 'Thu',
      active: false,
      coworkers: 2,
    },
    {
      number: 24,
      day: 'Fri',
      active: false,
      coworkers: 0,
    },
    {
      number: 25,
      day: 'Sat',
      active: false,
      coworkers: 0,
    },
    {
      number: 26,
      day: 'Sun',
      active: false,
      coworkers: 0,
    },
  
  ]
  return (
    <View
      style={styles.container}
    >
      {
        days.map((day, index) => {
          return (
            <CalendarDay
              key={index}
              number={day.number}
              day={day.day}
              active={day.active}
              coworkers={day.coworkers}
            />
          )
        })
      }
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    paddingTop: 23,
    flexDirection: 'row',
    flexWrap: 'wrap',
    gap: 8,
    rowGap: 16,
  },
});