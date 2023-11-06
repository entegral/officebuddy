import {
  StyleSheet,
  TouchableOpacity,
  Text,
  View
} from "react-native";

export default function GetStarted({
  headerBase,
  buttonBase,
  buttonText,
  textBase,
  linkBase,
}) {
  return (
    <View
      style={styles.mainDiv}
    >
      <View
        style={styles.contentsDiv}
      >
        <Text
          style={headerBase}
        >
          Workbuddies
        </Text>
      </View>

      <View
        styles={styles.buttonsDiv}
      >
        <TouchableOpacity
          style={buttonBase}
        >
          <Text
            style={buttonText}
          >
            GET STARTED
          </Text>
        </TouchableOpacity>
        <Text
          style={styles.buttonsDivText}
        >
          Already have an account? <Text style={linkBase}>Log In</Text>
        </Text>
        <Text
          style={{
            ...styles.buttonsDivText,
            fontWeight: '100',
            fontSize: 12,
          }}
        >
          By continuing, you agree to Workbuddies Terms of Service & Privacy Policy
        </Text>
      </View>

    </View>
  )
}

const styles = StyleSheet.create({
  mainDiv: {
    flex: 1,
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    gap: 24,
  },
  contentsDiv: {
    display: 'flex',
    paddingTop: 16,
    paddingBottom: 16,
  },
  buttonsDiv: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    flex: 1,
    gap: 12,
  },
  buttonsDivText: {
    color: '#fff',
    textAlign: 'center',
  }
});