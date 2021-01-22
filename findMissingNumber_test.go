package main

import "testing"

func TestFindMissingNumber( t *testing.T ){
	
	actual := testFindMissingNumber("23242526272830")
	
	if actual != "29" {
		t.Errorf("[ 23242526272830 ] failed, expected %v, got %v", "29", actual)
	}else{
		t.Logf("[ 23242526272830 ] success, expected %v, got %v", "29", actual)
	}

}

func TestFindMissingNumberTwo( t *testing.T ){
	
	actual := testFindMissingNumber("101102103104105106107108109111112113")
	
	if actual != "110" {
		t.Errorf("[ 101102103104105106107108109111112113 ] failed, expected %v, got %v", "110", actual)
	}else{
		t.Logf("[ 101102103104105106107108109111112113 ] success, expected %v, got %v", "110", actual)
	}

}

func TestFindMissingNumberThree( t *testing.T ){
	
	actual := testFindMissingNumber("12346789")
	
	if actual != "5" {
		t.Errorf("[ 12346789 ] failed, expected %v, got %v", "5", actual)
	}else{
		t.Logf("[ 12346789 ] success, expected %v, got %v", "5", actual)
	}

}

func TestFindMissingNumberFour( t *testing.T ){
	
	actual := testFindMissingNumber("79101112")
	
	if actual != "8" {
		t.Errorf("[ 79101112 ] failed, expected %v, got %v", "8", actual)
	}else{
		t.Logf("[ 79101112 ] success, expected %v, got %v", "8", actual)
	}

}

func TestFindMissingNumberFive( t *testing.T ){
	
	actual := testFindMissingNumber("7891012")
	
	if actual != "11" {
		t.Errorf("[ 7891012 ] failed, expected %v, got %v", "11", actual)
	}else{
		t.Logf("[ 7891012 ] success, expected %v, got %v", "11", actual)
	}

}

func TestFindMissingNumberSix( t *testing.T ){
	
	actual := testFindMissingNumber("9799100101102")
	
	if actual != "98" {
		t.Errorf("[ 9799100101102 ] failed, expected %v, got %v", "98", actual)
	}else{
		t.Logf("[ 9799100101102 ] success, expected %v, got %v", "98", actual)
	}

}

func TestFindMissingNumberSeven( t *testing.T ){
	
	actual := testFindMissingNumber("100001100002100003100004100006")
	
	if actual != "100005" {
		t.Errorf("[ 100001100002100003100004100006 ] failed, expected %v, got %v", "100005", actual)
	}else{
		t.Logf("[ 100001100002100003100004100006 ] success, expected %v, got %v", "100005", actual)
	}

}
