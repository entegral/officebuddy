import { StatusBar } from 'expo-status-bar';
import { StyleSheet, View } from 'react-native';
import LogInType from './LogIn/LogInType';

export default function App() {
  return (
    <View style={styles.container}>
      {/* <Text>Hello world!</Text> */}
      <LogInType
        headerBase={styles.headerBase}
        buttonBase={styles.buttonBase}
        buttonText={styles.buttonText}
      />
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
    //center text in div
    textAlign: 'center',
    
  },
  buttonBase: {
    borderColor: '#fff',
    borderRadius: 24,
    borderWidth: 1,
    backgroundColor: '#4B4949',
    color: '#fff',
    height: 50,
    width: 200,
  },
  buttonText: {
    textAlign: 'center',
    color: '#fff',
  }
  
});
