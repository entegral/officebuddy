import { View, Text, StyleSheet } from 'react-native';
import CalendarDay from './CalendarDay';

export default function Calendar({}) {
  const days =[
    {
      number: 20,
      day: 'Mon',
    },
    {
      number: 21,
      day: 'Tue',
    },
    {
      number: 22,
      day: 'Wed',
    },
    {
      number: 23,
      day: 'Thu',
    },
    {
      number: 24,
      day: 'Fri',
    },
    {
      number: 25,
      day: 'Sat',
    },
    {
      number: 26,
      day: 'Sun',
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