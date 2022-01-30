var board, zx, zy, clicks, possibles, clickCounter, oldzx = -1, oldzy = -1;
var number, secondsLimit, elapsedSeconds = 0, timeLabel, problemLabel;
var startButton, timerId, startToggle = 1, startTime, finished = 0;  
function getPossibles() {
    var ii, jj, cx = [-1, 0, 1, 0], cy = [0, -1, 0, 1];
    possibles = [];
    for( var i = 0; i < 4; i++ ) {
        ii = zx + cx[i]; jj = zy + cy[i];
        if( ii < 0 || ii > 3 || jj < 0 || jj > 3 ) continue;
        possibles.push( { x: ii, y: jj } );
    }
}
function updateBtns() {
    var b, v, id;
    for( var j = 0; j < 4; j++ ) {
        for( var i = 0; i < 4; i++ ) {
            id = "btn" + ( i + j * 4 );
            b = document.getElementById( id );
            v = board[i][j];
            if( v < 16 ) {
                b.innerHTML = ( "" + v );
                b.className = "button"
            }
            else {
                b.innerHTML = ( "" );
                b.className = "empty";
            }
        }
    }
    clickCounter.innerHTML = "Clicks: " + clicks;
}
function shuffleNumber() {
    number = 1 + (Math.floor( Math.random() * 2))
}
function goalBoard() {
    for( var j = 0; j < 4; j++ ) {
        for( var i = 0; i < 4; i++ ) {
            board[i][j] = 1 + i + (j * 4)
        }
    }
}
function boardOne() {
    board[0][0] = 2; 
    board[1][0] = 3; 
    board[2][0] = 4; 
    board[3][0] = 8; 
    board[0][1] = 6; 
    board[1][1] = 7;     
    board[2][1] = 12;     
    board[3][1] = 16; zx = 3; zy = 1;     
    board[0][2] = 1;     
    board[1][2] = 5;     
    board[2][2] = 10;     
    board[3][2] = 11;     
    board[0][3] = 9;         
    board[1][3] = 13;         
    board[2][3] = 14;         
    board[3][3] = 15;         
}
function boardTwo() {
    for( var j = 0; j < 2; j++ ) {
        for( var i = 0; i < 4; i++ ) {
            board[i][j] = 1 + (2*i) + (j * 8)
        }
    }
    for( var j = 2; j < 4; j++ ) {
        for( var i = 0; i < 4; i++ ) {
            board[i][j] = 2 + (2*i) + ((j-2) * 8)
        }
    }

}
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
    const response = new XMLHttpRequest();

    var json = JSON.stringify({
        problem: "A",
        numberA: number,
        elapsedSeconds: elapsedSeconds,
        movesA: clicks,
        complete: finished,
        sourceAddress: "",
        startTime: startTime
    });
    // {"problem": "A", "numberA": 1, "elapsedSeconds": 123, "movesA": 25, "sourceAddress": "1.2.3.4", "startTime": 1640975680}
    response.open("POST", 'http://127.0.0.1:8080/add')
    response.setRequestHeader('Content-Type', 'application/json');
    response.setRequestHeader('Accept', 'application/json');    
    console.log(json)
    response.send(json);

    response.onload = (e) => {
        alert(response.response);
    }
}
function stop() {
    clearInterval(timerId);
    sendData(); 
    toggleStartButton();
    restart();
}
function toggleStartButton() {
    if (startToggle) {
        startButton.innerHTML = "Stop";
        startButton.className = "stop";
        startButton.removeEventListener("click", startHandle);
        startButton.addEventListener( "click", stop, false );
        timeLabel.innerHTML = "Seconds left: " + secondsLimit;
        timerId = setInterval(countdown, 1000);
        startToggle = 0; 
        startTime = Math.round(new Date().getTime()/1000)  
    } else {
        startButton.innerHTML = "Start";
        startButton.className = "start";
        startButton.removeEventListener("click", stop);
        startButton.addEventListener( "click", startHandle, false );
        timeLabel.innerHTML = "Goal:";
        // timerId = setInterval(countdown, 1000);
        startToggle = 1; 
    }
}


function startHandle() {
    if (number == 1) { 
        boardOne(); 
    } else {
        boardTwo(); 
    }
    clicks = 0;
    updateBtns();
    toggleStartButton();
    // countdown();
}

function restart() {
    goalBoard();
    // shuffle();
    clicks = 0;
    secondsLimit = 300; 
    elapsedSeconds = 0;
    finished = 0; 
    updateBtns();
}
function checkFinished() {
    var a = 0;
    for( var j = 0; j < 4; j++ ) {
        for( var i = 0; i < 4; i++ ) {
            if( board[i][j] < a ) return false;
            a = board[i][j];
        }
    }
    finished = 1;
    return true;
}
function btnHandle( e ) {
    getPossibles();
    var c = e.target.i, r = e.target.j, p = -1;
    for( var i = 0; i < possibles.length; i++ ) {
        if( possibles[i].x == c && possibles[i].y == r ) {
            p = i;
            break;
        }
    }
    if( p > -1 ) {
        clicks++;
        var t = possibles[p];
        board[zx][zy] = board[t.x][t.y];
        zx = t.x; zy = t.y;
        board[zx][zy] = 16;
        updateBtns();
        if( checkFinished() ) {
            setTimeout(function(){ 
                alert( "WELL DONE!" );
                stop(); 
                // restart();
            }, 1);
        }
    }
}
function createBoard() {
    board = new Array( 4 );
    for( var i = 0; i < 4; i++ ) {
        board[i] = new Array( 4 );
    }
    for( var j = 0; j < 4; j++ ) {
        for( var i = 0; i < 4; i++ ) {
            board[i][j] = ( i + j * 4 ) + 1;
        }
    }
    zx = zy = 3; board[zx][zy] = 16;
}
function createSeries() {
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
    problemStatement.innerHTML = "Enter the next 8 letters in the sequence in the right row: ";
    document.body.appendChild( problemStatement );


    var b, d  = document.createElement( "div" );
    d.className += "board";
    document.body.appendChild( d );
    for( var i = 1; i < 3; i++ ) {
        for( var j = 1; j < 17; j++ ) {
            if (j < 9) {
                b = document.createElement( "button" ); // createElement
                // b.className += "btnNumber";
                b.id = "btn" + i + "-" + j ;
                b.className = "button"
                // b.i = i; b.j = j;
                // b.addEventListener( "click", btnHandle, false );
                b.appendChild( document.createTextNode( "" ) );
                d.appendChild( b );
            } else {
                b = document.createElement("input"); //input element, text
                b.id = "in" + i + "-" + j ;
                b.className = "button"
                b.setAttribute('type',"text");
                b.setAttribute('name',b.id);
                b.setAttribute('maxlength',"1");
                d.appendChild( b );
            }
        }
    }
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
    // id = "in1-10" ;
    // b = document.getElementById( id );
    // console.log("b: "+b.value)
    // "I-KLMN---J----OP"
}

