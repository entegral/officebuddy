import {
  View,
  Text,
  StyleSheet,
} from "react-native";
import Calendar from "./Calendar/Calendar";

export default function HomePage({

}){
  return (
    <View
      style={styles.container}
    >
      <Text
        style={styles.mainText}
      >
        SELECT YOUR AVAILABILITY
      </Text>
      <Calendar/>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    paddingTop: 48,
  },
  mainText: {
    color: '#fff',
    fontWeight: 700,
  }
});