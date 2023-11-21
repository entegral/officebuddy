import {
  View,
  Text,
  StyleSheet
} from 'react-native';
import CheckedSquareIcon from '../../Icons/CheckedSquare';

export default function CoworkerList({
  coworkersByDay,
}) {
  return (
    <View
      style={styles.container}
    >
      <Text style={styles.mainText}>
        🫂 WORKMATES PLANS
      </Text>
      {
        coworkersByDay.map((day, i) => {
          return (
            <View key={`${i}-${day}`}>
              <Text
                style={styles.dayText}
              >
                {day.day.toUpperCase()}
              </Text>
              <View>
                {day.coworkers.map((coworker, i) => {
                  return (
                    <View
                      key={coworker.id}
                    >
                      <Text
                        style={styles.coworkerText}
                      >
                       <CheckedSquareIcon
                        height={12}
                        width={12}  
                        />
                        &nbsp;&nbsp;{coworker.name}
                      </Text>
                    </View>
                  )
                })}
              </View>
            </View>
          )
        })
      }
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    marginTop: 27,
  },
  mainText: {
    color: '#fff',
    fontSize: 10,
  },
  dayText: {
    color: '#fff',
    fontSize: 14,
    fontWeight: 500,
    marginVertical: 14,
  },
  coworkerText:{
    color: '#fff',
    fontSize: 16,
  },
});