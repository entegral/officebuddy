import { View, Text, StyleSheet } from 'react-native';
import CalendarDay from './CalendarDay';

export default function Calendar({
  days,
  activeDays,
  activeDayHandler,
}) {
 
 
  const calculateCoWorkers = (day) => {
    let len = day.coworkers.length;
    if (activeDays[day.number]) {
      len++;
    }
    return len;
  }

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
              active={activeDays[day.number]}
              activeDayHandler={activeDayHandler}
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