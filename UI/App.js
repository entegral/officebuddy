import { StatusBar } from 'expo-status-bar';
import { SafeAreaView, StyleSheet, View } from 'react-native';
import LogInType from './LogIn/LogInType';
import GetStarted from './GetStarted/GetStarted';
import HomePage from './HomePage/HomePage';
import ToolBar from './ToolBar/ToolBar';

export default function App() {
  return (
    <SafeAreaView style={styles.container}>
      <ToolBar 
        style={styles.toolBar}
      />
      <View
        style={styles.mainContent}
      >
        <HomePage/>
      </View>
      {/* <Text>Hello world!</Text> */}
      {/* <GetStarted 
        headerBase={styles.headerBase}
        buttonBase={styles.buttonBase}
        buttonText={styles.buttonText}
        textBase={styles.textBase}
        linkBase={styles.linkBase}
      /> */}
      {/* <LogInType
        headerBase={styles.headerBase}
        buttonBase={styles.buttonBase}
        buttonText={styles.buttonText}
      /> */}

      <StatusBar style="light" />
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    color: '#fff',
    backgroundColor: '#272525',
  },
  toolBar: {
    height: 51,
  },
  mainContent:{
    padding: 21,
    flexGrow: 1,
  },
  // needs to be refactored?
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

//todo: 
//figure out how to add fonts
//add linter