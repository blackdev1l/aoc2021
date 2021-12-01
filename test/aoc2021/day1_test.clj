(ns aoc2021.day1-test
  (:require [clojure.test :refer :all]
            [aoc2021.day1 :refer :all]))


(def example '(199
               200
               208
               210
               200
               207
               240
               269
               260
               263))
(deftest day1
  (testing "exmaple"
    (call example)))
