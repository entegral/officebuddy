import { View, Text, StyleSheet } from 'react-native';
import CalendarDay from './CalendarDay';

export default function Calendar({days}) {
 
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
              coworkers={day.coworkers.length}
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