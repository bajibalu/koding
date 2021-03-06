kd = require 'kd'
{ PropTypes } = React = require 'react'
generateClassName = require 'classnames'
pluralize = require 'pluralize'
{ Grid, Row, Col } = require 'react-flexbox-grid'
dateDiffInDays = require 'app/util/dateDiffInDays'

classes = require './TrialChargeInfo.stylus'
textStyles = require 'lab/Text/Text.stylus'

Box = require 'lab/Box'
Label = require 'lab/Text/Label'

ExpirationMessage = ({ daysLeft }) ->

  daysLeft ?= 100

  title = if daysLeft >= 0
  then "Your trial period is about to expire"
  else "Your trial period has expired"

  <Row>
    <Col xs={12}>
      <Label size="small" type="danger">
        <strong>{title}</strong>
      </Label>
    </Col>
    <Col xs={12}>
      <Label size="small" type="info">
        Please add a credit card below to continue to use Koding. You will be
        charged based on your team size. If you have free credit,  it will be
        deducted from your bill after you enter a credit card. Click below for
        pricing details.
      </Label>
    </Col>
  </Row>


TrialChargeInfo = ({ teamSize, pricePerSeat, daysLeft }) ->

  isDanger = daysLeft < 4

  <Box border={1} type={if isDanger then 'danger' else 'secondary'}>
    <Row>
      {if daysLeft < 4
        <Col xs={12} className={classes['expiration']}>
          <ExpirationMessage daysLeft={daysLeft} />
        </Col>
      }
      <Col xs={5}>
        <Label size="small">
          Team Size: <strong>{pluralize 'Developer', teamSize, yes}</strong>
        </Label>
      </Col>
      <Col xs={7} className={textStyles.right}>
        <Label size="small">
          Monthly charge after trial: <strong>${pricePerSeat}</strong> per user
        </Label>
      </Col>
    </Row>
  </Box>


TrialChargeInfo.propTypes =
  teamSize: PropTypes.number
  pricePerSeat: PropTypes.number

TrialChargeInfo.defaultProps =
  teamSize: 1
  pricePerSeat: 49.97

module.exports = TrialChargeInfo

