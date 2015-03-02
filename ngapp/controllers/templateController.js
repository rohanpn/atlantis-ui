


app = angular.module('weather',['ui-router'])

app.factory('weathersvc', function(){
	var svc={};

	svc.getHourlyForecast = function(){
		var forecast=[];
		forecast.push({hour:'8am', temp: 50});
		return forecast;
	};

	svc.getWeeklyForecast = function(){
		var forecast= [];
		forecast.push({day :'monday',high :75; low:25});
		return forecast;
	};

	return svc;
});

app.controller('weatherctrl', function($scope, weathersvc){
	$scope.hourForecast= weathersvc.getHourlyForecast();
	$scope.weeklyForecast = weathersvc.getWeeklyForecast();

});


//////////////////////////////


app.config(function ($stateProvider, $urlRouterProvider){
	$urlRouterProvider.otherwise('/daily');

	$stateProvider
	.state('hourly',{
		url : '/hourly',
		controller : 'weatherctrl',
		templateUrl : 'partials/hourly.html'

	})
	.state('weekly',{

		url :'/weekly'
		controller : 'weatherctrl',
		templateUrl:'template/index2.html'

	})


});
