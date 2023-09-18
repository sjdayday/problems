var number, secondsLimit, elapsedSeconds = 0, timeLabel, problemLabel;
var startButton, timerId, startToggle = 1, startTime, notFinished = 1;  
var answerLabel, notYet = 0, attemptsB = 0; 
function countdown() {
    secondsLimit--;
    elapsedSeconds++; 
    if (secondsLimit > 0) {
        timeLabel.innerHTML = "Seconds left: " + secondsLimit;   
    } else {
        stop()
    }  
}
function sendData() {
    if (notFinished) {
        const response = new XMLHttpRequest();

        var json = JSON.stringify({
            problem: "B",
            elapsedSeconds: elapsedSeconds,
            complete: 1,
            attemptsB: attemptsB,
            sourceAddress: "",
            startTime: startTime
        });
        // {"problem": "A", "numberA": 1, "elapsedSeconds": 123, "movesA": 25, "sourceAddress": "1.2.3.4", "startTime": 1640975680}
        response.open("POST", 'http://127.0.0.1:80/add')
        response.setRequestHeader('Content-Type', 'application/json');
        response.setRequestHeader('Accept', 'application/json');    
        console.log(json)
        response.send(json);
        notFinished = 0; 
        response.onload = (e) => {
            alert(response.response);
        }
    }
}
function sendAnswer( answer ) {
    const response = new XMLHttpRequest();

    var json = JSON.stringify({
        answer: answer 
    });
    response.open("POST", 'http://127.0.0.1:80/check')
    response.setRequestHeader('Content-Type', 'application/json');
    response.setRequestHeader('Accept', 'application/json'); 
    response.onreadystatechange = function () {
        if (response.readyState === 4 && response.status === 200) { 
            answerLabel.innerHTML = "You got it!";
            attemptsB++;
            stop();
        } else if (response.readyState === 4 && response.status === 400) {
            attemptsB++;
            if (notYet) {
                answerLabel.innerHTML = "Please keep trying";
                notYet = 0; 
            } else {
                answerLabel.innerHTML = "Not yet...try again";
                notYet = 1; 
            }
        } else {
            console.log("ready: "+response.readyState+" status: "+response.status);
        }
    };   
    // console.log(json);
    response.send(json);
}

function stop() {
    clearInterval(timerId);
    sendData(); 
}

function createSeries() {
    answerLabel = document.createElement( "p" );
    answerLabel.className += "txt";
    answerLabel.innerHTML = "";
    document.body.appendChild( answerLabel );

    problemLabel = document.createElement( "p" );
    problemLabel.className += "txt";
    problemLabel.innerHTML = "Problem B";
    document.body.appendChild( problemLabel );

    timeLabel = document.createElement( "p" );
    timeLabel.className += "txt";
    timeLabel.innerHTML = "Seconds left: " + secondsLimit;
    document.body.appendChild( timeLabel );

    problemStatement = document.createElement( "p" );
    problemStatement.className += "txt";
    problemStatement.innerHTML = "Enter the next 8 letters in the right row: ";
    document.body.appendChild( problemStatement );

    var tbl = document.createElement('table');
    var b  = document.createElement( "div" );
    var tr, td; 
    tbl.className += "board";
    
    for( var i = 1; i < 3; i++ ) {
        tr = tbl.insertRow();
        tr.className = "row";
        for( var j = 1; j < 17; j++ ) {
            td = tr.insertCell();
            if (j < 9) {
                b = document.createElement( "button" ); 
                b.id = "btn" + i + "-" + j ;
                b.className = "button"
                b.appendChild( document.createTextNode( "" ) );
                td.appendChild( b );
            } else {
                b = document.createElement("input"); 
                b.id = "in" + i + "-" + j ;
                b.className = "button"
                b.setAttribute('type',"text");
                b.setAttribute('name',b.id);
                b.setAttribute('maxlength',"1");
                td.appendChild( b );
            }
        }
    }
    document.body.appendChild( tbl );
    setLetter(1,1,"A");
    setLetter(2,2,"B");
    setLetter(2,3,"C");
    setLetter(2,4,"D");
    setLetter(1,5,"E");
    setLetter(1,6,"F");
    setLetter(2,7,"G");
    setLetter(1,8,"H");
 

    var sbtn  = document.createElement( "div" );
    sbtn.className += "startd";
    document.body.appendChild( sbtn );
    startButton = document.createElement( "button" );
    startButton.className += "start";
    startButton.addEventListener( "click", checkSolution, false );
    startButton.innerHTML = "Check solution";
 
    sbtn.appendChild( startButton );
    clickCounter = document.createElement( "p" );
    clickCounter.className += "txt";
    document.body.appendChild( clickCounter );

    timerId = setInterval(countdown, 1000);
}
function start() {
    secondsLimit = 300; 
    elapsedSeconds = 0;
    finished = 0; 
    startToggle = 0; 
    startTime = Math.round(new Date().getTime()/1000) 
    createSeries();
}
function setLetter(i, j, letter) {
    var b, id;
    id = "btn" + i + "-" + j;
    console.log(id)
    b = document.getElementById( id );
    b.innerHTML = ( ""+letter );
    b.className = "button"
}
function checkSolution() {
    if (notFinished) {
        var a, b, id, answer;
        console.log("check solution called");
        answer = "";
        for ( var i = 0; i < 16; i++ ) {
            a = document.getElementsByTagName("input")[i].value;
            if (a.trim() == "") {
                answer = answer + "-";
            } else {
                answer = answer + a.toUpperCase();
            }
        }
        console.log("answer: "+answer )
        sendAnswer( answer );
    }
}
