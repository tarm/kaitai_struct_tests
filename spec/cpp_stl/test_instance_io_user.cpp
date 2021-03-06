#include <boost/test/unit_test.hpp>

#include <instance_io_user.h>

#include <iostream>
#include <fstream>
#include <vector>

BOOST_AUTO_TEST_CASE(test_instance_io_user) {
    std::ifstream ifs("src/instance_io.bin", std::ifstream::binary);
    kaitai::kstream ks(&ifs);
    instance_io_user_t* r = new instance_io_user_t(&ks);

    BOOST_CHECK_EQUAL(r->qty_entries(), 3);

    BOOST_CHECK_EQUAL(r->entries()->at(0)->name(), std::string("the"));
    BOOST_CHECK_EQUAL(r->entries()->at(1)->name(), std::string("rainy"));
    BOOST_CHECK_EQUAL(r->entries()->at(2)->name(), std::string("day it is"));

    delete r;
}
