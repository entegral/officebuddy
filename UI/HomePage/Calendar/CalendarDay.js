import {
  View,
  Text,
  StyleSheet,
} from "react-native";

export default function CalendarDay({
  number,
  day,
}) {
  return (
    <View
      style={styles.container}
    >
      <Text
        style={styles.numberText}
      >
        {number}
      </Text>
      <Text
        style={styles.dayText}
      >
        {day}
      </Text>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    display: 'flex',
    width: 64,
    height: 80,
    paddingVertical: 12,
    paddingHorizontal: 8,
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    borderColor: '#CDCDD0',
    borderRadius: 24,
    borderWidth: 2,
    backgroundColor: '#4B4949',
  },
  numberText: {
    fontSize: 20,
    fontWeight: 500,
    color: '#CDCDD0',
  },
  dayText: {
    fontSize: 12,
    color: '#CDCDD0',
  }
});