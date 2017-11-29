'use strict';

/**
 * @ngdoc function
 * @name angularjsHasuraApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the angularjsHasuraApp
 */
angular.module('angularjsHasuraApp')
  .controller('MainCtrl', function ($http, $window) {
    this.sampleText = "Hello from AngularJS deployed on Hasura";

    // // Hasura Data API Example
    // var url = "https://data.[cluster-name].hasura-app.io/v1/query";
    // var token = $window.localStorage["token"];
    // var config = {
    //   headers:{
    //     "content-type": "application/json",
    //     "authorization": "Bearer " + token
    //   }
    // };
    // var payload = {
    //   type: "insert",
    //   args: {
    //     table: "todo",
    //     "objects": [
    //       {title: "Buy groceries", category: "personal"}
    //     ],
    //     returning: ["id"]
    //   }
    // };

    // $http.post(url, payload, config).then(function(response){
    //   console.log(response.data);
    // }, function(response){
    //   console.log(response.data);
    // });
  });
