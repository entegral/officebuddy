import {
  View,
  Text,
  StyleSheet,
} from "react-native";

export default function CalendarDay({
  number,
  day,
  active = false,
  coworkers,
}) {

  const coworkerPips = () => {
    const colors = ['#FDAFFF', '#AFCAFF', '#B1FFAF']
    const pips = [];
    const length = coworkers > 3 ? 3 : coworkers;
    for (let i = 0; i < length; i++) {
      pips.push(
        <View
          style={[styles.pip, {right: i * 10, backgroundColor: colors[i]}]}
          key={i}
        >
          {i === (length - 1) && <Text style={styles.pipText}>{coworkers}</Text>}
        </View>
      )
    }
    return pips;
  };
  return (
    <View
      style={styles.day}
    >
      <View
        style={[styles.container, active && styles.containerActive]}
      >

        <Text
          style={[styles.numberText, active && styles.textActive]}
        >
          {number}
        </Text>
        <Text
          style={[styles.dayText, active && styles.textActive]}
        >
          {day}
        </Text>
      </View>
      <View style={styles.pipContainer}>
        {coworkerPips()}
      </View>
    </View>
    
  )
}

const styles = StyleSheet.create({
  day: {
    display: 'flex',
    width: 64,
    height: 82,
  },
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
    position: 'relative',
  },
  containerActive: {
    borderColor: '#D2FFAF',
    backgroundColor: '#D2FFAF',
  },
  pipContainer: {
    width: 64,
    height: 24,
    position: 'absolute',
    top: -12,
  },
  pip: {
    backgroundColor: 'white',
    height: 24,
    width: 24,
    borderRadius: 12,
    top: 0,
    position: 'absolute',
  },
  pipText: {
    //center text in parent
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: [{translateX: -5}, {translateY: -8}],
    fontWeight: 700,
  },
  numberText: {
    fontSize: 20,
    fontWeight: 500,
    color: '#CDCDD0',
  },
  dayText: {
    fontSize: 12,
    color: '#CDCDD0',
  },
  textActive: {
    color: '#000'
  }
});