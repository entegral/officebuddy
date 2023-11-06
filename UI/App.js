import { StatusBar } from 'expo-status-bar';
import { StyleSheet, View } from 'react-native';
import LogInType from './LogIn/LogInType';
import GetStarted from './GetStarted/GetStarted';

export default function App() {
  return (
    <View style={styles.container}>
      {/* <Text>Hello world!</Text> */}
      <GetStarted 
        headerBase={styles.headerBase}
        buttonBase={styles.buttonBase}
        buttonText={styles.buttonText}
        textBase={styles.textBase}
        linkBase={styles.linkBase}
      />
      {/* <LogInType
        headerBase={styles.headerBase}
        buttonBase={styles.buttonBase}
        buttonText={styles.buttonText}
      /> */}
      <StatusBar style="auto" />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#000',
    color: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
  headerBase: {
    color: '#fff',
    fontSize: 34,
    fontWeight: 400,
    textAlign: 'center'
  },
  buttonBase: {
    borderColor: '#fff',
    borderRadius: 24,
    borderWidth: 1,
    backgroundColor: '#4B4949',
    color: '#fff',
    height: 50,
    alignSelf: 'stretch',
    paddingTop: 16,
    paddingBottom: 16,
    paddingLeft: 20,
    paddingRight: 20,
    marginBottom: 12,
  },
  buttonText: {
    textAlign: 'center',
    color: '#fff',
    fontSize: 14,

  },
  textBase: {
    color: '#fff'
  },
  linkBase: {
    color: '#D2FFAF',
    fontWeight: 700,
  }
});
